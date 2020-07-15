<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="订单id">
          <el-input placeholder="搜索条件" v-model="searchInfo.id" clearable></el-input>
        </el-form-item>    
        <el-form-item label="订单号">
          <el-input placeholder="搜索条件" v-model="searchInfo.out_trade_no" clearable></el-input>
        </el-form-item>    
        <el-form-item label="用户id">
          <el-input placeholder="搜索条件" v-model="searchInfo.user_id" clearable></el-input>
        </el-form-item>    
        <el-form-item label="用户电话">
          <el-input placeholder="搜索条件" v-model="searchInfo.mobile" clearable></el-input>
        </el-form-item>    
        <el-form-item label="微信open_id">
          <el-input placeholder="搜索条件" v-model="searchInfo.open_id" clearable></el-input>
        </el-form-item>        
        <el-form-item label="订单状态">
          <el-select v-model="searchInfo.status" placeholder="请选择订单状态搜索" clearable :style="{width: '100%'}">
            <el-option v-for="(item, index) in statusOptions" :key="index" :label="item.label"
                       :value="item.value" :disabled="item.disabled"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
<!--        <el-form-item>-->
<!--          <el-button @click="openDialog" type="primary">新增订单</el-button>-->
<!--        </el-form-item>-->
      </el-form>
    </div>
    <el-button @click="clearFilter">清除所有过滤器</el-button>
    <el-table
      :data="tableData"
      border
      ref="multipleTable"
      stripe
      style="width: 100%"
      tooltip-effect="dark"
    >
    <el-table-column type="selection" width="55"></el-table-column>

    <el-table-column label="订单id" prop="id" width="120"></el-table-column> 
    
    <el-table-column label="订单号" prop="out_trade_no" width="120"></el-table-column> 
    
    <el-table-column label="用户id" prop="user_id" width="120"></el-table-column>
    
    <el-table-column label="用户电话" prop="mobile" width="120"></el-table-column>
    
    <el-table-column label="微信open_id" prop="open_id" width="120"></el-table-column> 
    
    <el-table-column label="交易总金额" prop="amount" width="120">
      <template slot-scope="scope">{{scope.row.amount|formatPrice}}</template>
    </el-table-column>
    
    <el-table-column label="退款时间" prop="refund_time" width="120">
      <template slot-scope="scope">{{scope.row.refund_time|formatDate}}</template>
    </el-table-column>
    
    <el-table-column label="订单状态"
                     prop="status"
                     width="120"
                     :filters="[{text: '待支付', value: 0}, {text: '付款成功', value: 2}, {text: '已退款', value: 3}, {text: '已过期/关闭', value: 4}, {text: '体检完毕，待评价', value: 5}]"
                     :filter-method="filterStatus"
                     sortable
                     filter-placement="bottom-end">
      <template slot-scope="scope">

        <el-tag
                :type="scope.row.status === 0 ? 'primary' : scope.row.status === 2 ? 'success' : scope.row.status === 3 ? 'warning' : 'info'"
                disable-transitions> {{scope.row.status|formatOrderStatus}}
        </el-tag>
      </template>
    </el-table-column>

    <el-table-column label="订单详细项目"  width="140">
      <template slot-scope="scope">
        <el-button @click="ShowOrderItem(scope.row)" size="small" type="primary">查看订单详细项目</el-button>
      </template>
    </el-table-column>
    
    <el-table-column label="订单备注" prop="remark" width="120"></el-table-column> 
    
    <el-table-column label="创建时间" prop="create_time" width="120">
      <template slot-scope="scope">{{scope.row.create_time|formatDate}}</template>
    </el-table-column>
    
    <el-table-column label="更新时间" prop="update_time" width="120">
      <template slot-scope="scope">{{scope.row.update_time|formatDate}}</template>
    </el-table-column>
    
      <el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-button @click="updateOrder(scope.row)" size="small" type="primary">变更</el-button>
          <el-popover placement="top" width="160" v-model="scope.row.visible">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin: 0">
              <el-button size="mini" type="text" @click="scope.row.visible = false">取消</el-button>
              <el-button type="primary" size="mini" @click="deleteOrder(scope.row)">确定</el-button>
            </div>
            <el-button type="danger" icon="el-icon-delete" size="mini" slot="reference">删除</el-button>
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

    <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="弹窗操作">
      <el-row :gutter="15">
        <el-form ref="elForm" :model="formData" :rules="rules" size="medium" label-width="100px">
          <el-col :span="7">
            <el-form-item label="订单id" prop="id">
              <el-input-number v-model="formData.id" placeholder="订单id" :disabled='true'></el-input-number>
            </el-form-item>
          </el-col>
          <el-col :span="7">
            <el-form-item label="订单状态" prop="status">
              <el-select v-model="formData.status" placeholder="请选择订单状态" clearable :style="{width: '100%'}">
                <el-option v-for="(item, index) in statusOptions" :key="index" :label="item.label"
                           :value="item.value" :disabled="item.disabled"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="24">
            <el-form-item label="备注" prop="remark">
              <el-input v-model="formData.remark" type="textarea" placeholder="请输入备注" :maxlength="500"
                        :autosize="{minRows: 6, maxRows: 6}" :style="{width: '100%'}"></el-input>
            </el-form-item>
          </el-col>
        </el-form>
      </el-row>
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button @click="enterDialog" type="primary">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import {
    createOrder,
    deleteOrder,
    updateOrder,
    findOrder,
    getOrderList
} from "@/api/order";  //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/data";
import infoList from "@/components/mixins/infoList";
import { formatPrice } from "@/utils/price";
let that
export default {
  name: "Order",
  mixins: [infoList],
  data() {
    return {
      listApi: getOrderList,
      dialogFormVisible: false,
      visible: false,
      type: "",
      formData: {
        id: null,
        status: null,
        remark: null,
      },
      rules: {
        id: [{
          required: true,
          message: '订单id',
          trigger: 'blur'
        }],
        status: [{
          required: true,
          message: '请选择订单状态',
          trigger: 'change'
        }],
        remark: [{
          message: '请输入备注',
          trigger: 'blur'
        }],
      },
      statusOptions: [{
        "label": "待支付",
        "value": 0
      }, {
        "label": "付款成功",
        "value": 2
      }, {
        "label": "已退款",
        "value": 3
      }, {
        "label": "已过期/关闭",
        "value": 4
      }, {
        "label": "体检完毕/待评价",
        "value": 5
      }],
      statusDict: new Map([[0,"待支付"], [2, "付款成功"], [3, "已退款"],
        [4, "已过期/关闭"], [5, "体检完毕/待评价"]])
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
    formatBoolean: function(bool) {
      if (bool != null) {
        return bool ? "是" :"否";
      } else {
        return "";
      }
    },
    formatOrderStatus: function(orderStatus) {
      return that.statusDict.get(orderStatus);
    },
    formatPrice: formatPrice
  },
  methods: {
      //条件搜索前端看此方法
      onSubmit() {
        this.page = 1
        this.pageSize = 10               
        this.getTableData()
      },
    ShowOrderItem(row) {
        this.$router.push({name: "orderItem", params: {order_id: row.id}})
    },
    clearFilter() {
      this.$refs.multipleTable.clearFilter();
    },
    filterStatus(value, row) {
      return row.status === value;
    },
    async updateOrder(row) {
      const res = await findOrder({ id: row.id });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.reorder;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
        
          id:null,
          out_trade_no:null,
          user_id:null,
          mobile:null,
          open_id:null,
          amount:null,
          refund_time:null,
          status:null,
          remark:null,
          create_time:null,
          update_time:null,
      };
    },
    async deleteOrder(row) {
      this.visible = false;
      const res = await deleteOrder({ id: row.id });
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "删除成功"
        });
        await this.getTableData();
      }
    },
    async enterDialog() {
      let res;
      switch (this.type) {
        case "create":
          res = await createOrder(this.formData);
          break;
        case "update":
          res = await updateOrder(this.formData);
          break;
        default:
          res = await createOrder(this.formData);
          break;
      }
      if (res.code == 0) {
        this.$message({
          type:"success",
          message:"创建/更改成功"
        })
        this.closeDialog();
        await this.getTableData();
      }
    },
    openDialog() {
      this.type = "create";
      this.dialogFormVisible = true;
    }
  },
  created() {
    this.getTableData();
  },
  beforeCreate() {
    that = this;
  }
};
</script>

<style>
</style>

