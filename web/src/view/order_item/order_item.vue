<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="id">
          <el-input placeholder="搜索条件" v-model="searchInfo.id" clearable></el-input>
        </el-form-item>    
        <el-form-item label="用户id">
          <el-input placeholder="搜索条件" v-model="searchInfo.user_id" clearable></el-input>
        </el-form-item>    
        <el-form-item label="订单id">
          <el-input placeholder="搜索条件" v-model="searchInfo.order_id" clearable></el-input>
        </el-form-item>    
        <el-form-item label="套餐id">
          <el-input placeholder="搜索条件" v-model="searchInfo.pkg_id" clearable></el-input>
        </el-form-item>      
        <el-form-item label="体检人姓名">
          <el-input placeholder="搜索条件" v-model="searchInfo.examinee_name" clearable></el-input>
        </el-form-item>    
        <el-form-item label="体检人电话">
          <el-input placeholder="搜索条件" v-model="searchInfo.examinee_mobile" clearable></el-input>
        </el-form-item>                
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
<!--        <el-form-item>-->
<!--          <el-button @click="openDialog" type="primary">新增订单项目</el-button>-->
<!--        </el-form-item>-->
      </el-form>
    </div>
    <el-table
      :data="tableData"
      border
      ref="multipleTable"
      stripe
      style="width: 100%"
      tooltip-effect="dark"
    >
    <el-table-column type="selection" width="55"></el-table-column>

    <el-table-column label="id" prop="id" width="120"></el-table-column> 
    
    <el-table-column label="用户id" prop="user_id" width="120"></el-table-column> 
    
    <el-table-column label="订单id" prop="order_id" width="120"></el-table-column>

    <el-table-column label="套餐id" prop="pkg_id" width="120">
      <template slot-scope="scope">
        <el-button @click="getPkgDetail(scope.row)" size="small" type="primary">{{scope.row.pkg_id}}</el-button>
      </template>
    </el-table-column>
    
    <el-table-column label="套餐的单价" prop="pkg_price" width="120">
      <template slot-scope="scope">{{scope.row.pkg_price|formatPrice}}</template>
    </el-table-column>
    
    <el-table-column label="体检人姓名" prop="examinee_name" width="120"></el-table-column> 
    
    <el-table-column label="体检人电话" prop="examinee_mobile" width="120"></el-table-column> 
    
    <el-table-column label="体检人身份证号码" prop="id_card_no" width="120"></el-table-column> 
    
    <el-table-column label="体检人是否已婚" prop="is_married" width="120">
      <template slot-scope="scope">{{scope.row.is_married|formatBoolean}}</template>
    </el-table-column>
    
    <el-table-column label="性别" prop="gender" width="120">
      <template slot-scope="scope">{{scope.row.gender|formatGender}}</template>
    </el-table-column>
    
    <el-table-column label="体检日期" prop="examine_date" width="120">

      <template slot-scope="scope">
        {{scope.row.examine_date|formatDay}}
      </template>

    </el-table-column>
    
    <el-table-column label="创建时间" prop="create_time" width="120">
      <template slot-scope="scope">{{scope.row.create_time|formatDate}}</template>

    </el-table-column>
    
    <el-table-column label="更新时间" prop="update_time" width="120">
      <template slot-scope="scope">{{scope.row.update_time|formatDate}}</template>

    </el-table-column>
    
    <el-table-column label="按钮组">
      <template slot-scope="scope">
        <el-button @click="notifyAppointmentOk(scope.row)" size="small" type="primary" >短信通知预约成功</el-button>
        <el-button @click="updateOrderItem(scope.row)" size="small" type="primary" >变更</el-button>
        <el-popover placement="top" width="160" v-model="scope.row.visible">
          <p>确定要删除吗？</p>
          <div style="text-align: right; margin: 0">
            <el-button size="mini" type="text" @click="scope.row.visible = false">取消</el-button>
            <el-button type="primary" size="mini" @click="deleteOrderItem(scope.row)">确定</el-button>
          </div>
          <el-button type="danger" icon="el-icon-delete" size="mini" slot="reference" disabled>删除</el-button>
        </el-popover>
      </template>
    </el-table-column>
  </el-table>

    <el-pagination
      :current-page="page"
      :page-size="pageSize"
      :page-sizes="[10, 30, 50, 100]"
      :style="{float:'right',padding:'20px'}"
      :total="total"
      @current-change="handleCurrentChange"
      @size-change="handleSizeChange"
      layout="total, sizes, prev, pager, next, jumper"
    ></el-pagination>

    <el-dialog :before-close="closePkgDialog" :visible.sync="dialogPkgFormVisible" title="弹窗操作">
      <el-row :gutter="15">
        <el-form ref="elForm" :model="pkgFormData" size="large" label-width="100px" label-height="100px">
          <el-col :span="17">
            <el-form-item label="套餐所属医院" prop="hospital_id">
              <el-select v-model="pkgFormData.hospital_id" placeholder="套餐所属医院" clearable disabled
                         :style="{width: '100%'}">
                <el-option v-for="(item, index) in hospital_idOptions" :key="index" :label="item.label"
                           :value="item.value" :disabled="item.disabled"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="17">
            <el-form-item label="套餐名" prop="name">
              <el-input v-model="pkgFormData.name" placeholder="套餐名" :maxlength="64" clearable disabled
                        :style="{width: '100%'}"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="17">
            <el-form-item label="套餐目标人群" prop="target">
              <el-select v-model="pkgFormData.target" placeholder="请选择套餐目标人群" clearable disabled :style="{width: '100%'}">
                <el-option v-for="(item, index) in targetOptions" :key="index" :label="item.label"
                           :value="item.value" :disabled="item.disabled"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="17">
            <el-form-item label="门市价" prop="price_original">
              <el-input-number v-model="pkgFormData.price_original" placeholder="门市价" disabled></el-input-number>
            </el-form-item>
          </el-col>
          <el-col :span="17">
            <el-form-item label="真实价格" prop="price_real">
              <el-input-number v-model="pkgFormData.price_real" placeholder="真实价格" disabled></el-input-number>
            </el-form-item>
          </el-col>
          <el-col :span="17">
            <el-form-item label="套餐概述" prop="brief">
              <el-input v-model="pkgFormData.brief" type="textarea" placeholder="请输入套餐概述" disabled
                        :autosize="{minRows: 4, maxRows: 4}" :style="{width: '100%'}"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="17">
            <el-form-item label="套餐备注" prop="comment">
              <el-input v-model="pkgFormData.comment" type="textarea" placeholder="请输入套餐备注" disabled
                        :autosize="{minRows: 4, maxRows: 4}" :style="{width: '100%'}"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="17">
            <el-form-item label="套餐温馨提示" prop="tips">
              <el-input v-model="pkgFormData.tips" type="textarea" placeholder="请输入套餐温馨提示" disabled
                        :autosize="{minRows: 4, maxRows: 4}" :style="{width: '100%'}"></el-input>
            </el-form-item>
          </el-col>
<!--          <el-col :span="24">-->
<!--            <el-form-item size="large">-->
<!--              <el-button @click="closePkgDialog">确定</el-button>-->
<!--            </el-form-item>-->
<!--          </el-col>-->

        </el-form>

      </el-row>
    </el-dialog>


      <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="弹窗操作">
<!--        此处请使用表单生成器生成form填充 表单默认绑定 formData 如手动修改过请自行修改key-->
        <el-form ref="elFormOrderItem" :model="formData" :rules="rules" size="medium" label-width="100px">
          <el-col :span="12">
            <el-form-item label="体检日期" prop="examine_date" >
              <el-date-picker v-model="formData.examine_date" format="yyyy-MM-dd" value-format="timestamp"
                              :style="{width: '100%'}" placeholder="请选择体检日期" clearable></el-date-picker>
            </el-form-item>
          </el-col>
          <el-col :span="24">
            <el-form-item label="体检人姓名" prop="examinee_name">
              <el-input v-model="formData.examinee_name" placeholder="请输入体检人姓名" clearable
                        :style="{width: '100%'}"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="24">
            <el-form-item label="体检人身份证号码" prop="id_card_no">
              <el-input v-model="formData.id_card_no" placeholder="请输入体检人身份证" clearable
                        :style="{width: '100%'}"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="24">
            <el-form-item label="体检人电话" prop="examinee_mobile">
              <el-input v-model="formData.examinee_mobile" placeholder="请输入体检人电话" clearable
                        :style="{width: '100%'}"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="24">
            <el-form-item size="large">
              <el-button @click="resetForm">重置</el-button>
            </el-form-item>
          </el-col>
        </el-form>

        <div class="dialog-footer" slot="footer">
          <el-button @click="closeDialog">取 消</el-button>
          <el-button @click="enterDialog" type="primary">确 定</el-button>
        </div>
      </el-dialog>


  </div>
</template>

<script>
import {
    createOrderItem,
    deleteOrderItem,
    updateOrderItem,
    findOrderItem,
    getOrderItemList
} from "@/api/order_item";  //  此处请自行替换地址
import {notifyAppointmentOk} from "@/api/notify"
import { formatTimeToStr } from "@/utils/data";
import infoList from "@/components/mixins/infoList";
import {findPackage} from "../../api/pkg";
import {formatPrice} from "@/utils/price"
export default {
  name: "OrderItem",
  mixins: [infoList],
  data() {
    return {
      listApi: getOrderItemList,
      dialogFormVisible: false,
      dialogPkgFormVisible: false,
      visible: false,
      type: "",
      formData: {
        id:null,user_id:null,order_id:null,pkg_id:null,pkg_price:null,examinee_name:null,examinee_mobile:null,id_card_no:null,is_married:null,gender:null,examine_date:null,create_time:null,update_time:null,
      },
      pkgFormData: {
        id: null,
        hospital_id: null,
        name: null,
        target: null,
        disease: null,
        price_original: null,
        price_real: null,
        avatar_url: null,
        brief: null,
        comment: null,
        tips: null,
        create_time: null,
        update_time: null,
        ctg_ids: []
      },
      rules: {
        examinee_name: [],
        id_card_no: [{
          required: true,
          message: '请输入体检人身份证',
          trigger: 'blur'
        }],
        examinee_mobile: [{
          required: true,
          message: '请输入体检人电话',
          trigger: 'blur'
        }],
      },
      targetOptions: [{
        "label": "不限",
        "value": 0
      }, {
        "label": "男士",
        "value": 1
      }, {
        "label": "女-未婚",
        "value": 2
      }, {
        "label": "女-已婚",
        "value": 3
      }],
    };
  },
  filters: {
    formatDate: function(time) {
      if (time != null && time != "") {
        var date = new Date(time * 1000);
        return formatTimeToStr(date, "yyyy-MM-dd hh:mm:ss");
      } else {
        return "";
      }
    },
    formatDay: function(time) {
      if (time != null && time != "") {
        var date = new Date(time * 1000);
        return formatTimeToStr(date, "yyyy-MM-dd");
      } else {
        return "";
      }
    },
    formatPrice: formatPrice,
    formatGender: function(gender) {
      if (gender === 1) {
        return "男";
      } else if (gender === 2) {
        return "女";
      } else {
        return "未知"
      }

    },
    formatBoolean: function(bool) {
      if (bool != null) {
        return bool ? "是" :"否";
      } else {
        return "";
      }
    }
  },
  methods: {
      //条件搜索前端看此方法
      onSubmit() {
        this.page = 1
        this.pageSize = 10                 
        this.getTableData()
      },
    resetForm() {
      this.$refs['elFormOrderItem'].resetFields()
    },
    async getPkgDetail(row) {
      const res = await findPackage({id: row.pkg_id})
      if (res.code === 0) {
        this.pkgFormData = res.data.repkg;
        this.dialogPkgFormVisible=true;
      }

    },
    async getAllOrderItemByOrderId(page = this.page, pageSize = this.pageSize) {
      const order_id = this.$route.params.order_id
      const table = await this.listApi({page, pageSize, order_id})
      this.tableData = table.data.list
      this.total = table.data.total
      this.page = table.data.page
      this.pageSize = table.data.pageSize
    },
    async updateOrderItem(row) {
      const res = await findOrderItem({ id: row.id });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.reorderItem;
        this.formData.examine_date *= 1000
        this.dialogFormVisible = true;
      }
    },
    closePkgDialog() {
      this.dialogPkgFormVisible = false;
      this.pkgFormData = {
        id: null,
        hospital_id: null,
        name: null,
        target: null,
        disease: null,
        price_original: null,
        price_real: null,
        avatar_url: null,
        brief: null,
        comment: null,
        tips: null,
        create_time: null,
        update_time: null,
        ctg_ids: []
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
        
          id:null,
          user_id:null,
          order_id:null,
          pkg_id:null,
          pkg_price:null,
          examinee_name:null,
          examinee_mobile:null,
          id_card_no:null,
          is_married:null,
          gender:null,
          examine_date:null,
          create_time:null,
          update_time:null,
      };
    },
    async notifyAppointmentOk(row) {
      const res = await notifyAppointmentOk({
        out_trade_no: this.$route.query.out_trade_no,
        pkg_id: row.pkg_id,
        examinee_name: row.examinee_name,
        examine_date: row.examine_date,
        examinee_mobile: row.examinee_mobile
      });
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "发送成功"
        });
      }
    },
    async deleteOrderItem(row) {
      this.visible = false;
      const res = await deleteOrderItem({ id: row.id });
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "删除成功"
        });
        this.getAllOrderItemByOrderId(1, 1000);
      }
    },
    async enterDialog() {
      let res;
      this.formData.examine_date *= 0.001
      switch (this.type) {
        case "create":
          res = await createOrderItem(this.formData);
          break;
        case "update":
          res = await updateOrderItem(this.formData);
          break;
        default:
          res = await createOrderItem(this.formData);
          break;
      }
      if (res.code == 0) {
        this.$message({
          type:"success",
          message:"创建/更改成功"
        })
        this.closeDialog();
        await this.getAllOrderItemByOrderId(1, 1000);
      }
    },
    openDialog() {
      this.type = "create";
      this.dialogFormVisible = true;
    }
  },
  created() {
    this.getAllOrderItemByOrderId(1, 1000);
  }
};
</script>

<style>
</style>