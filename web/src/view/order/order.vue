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
        <el-form-item label="只看申请退款的订单">
          <el-checkbox placeholder="搜索条件" v-model="searchInfo.is_refund_applied" clearable></el-checkbox>
        </el-form-item>
        <el-form-item label="订单状态">
          <el-select v-model="searchInfo.status" placeholder="请选择订单状态搜索" clearable :style="{width: '100%'}">
            <el-option v-for="(item, index) in statusOptions" :key="index" :label="item.label"
                       :value="item.value" ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="跟进状态">
          <el-select v-model="searchInfo.handle_status_filter" placeholder="请选择订单跟进状态搜索" :style="{width: '100%'}" clearable>
            <el-option v-for="(item, index) in handleStatusOptions" :key="index" :label="item.label"
                       :value="item.value" ></el-option>
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

      <el-table-column label="订单详细项目" width="140">
        <template slot-scope="scope">
          <el-button @click="ShowOrderItem(scope.row)" size="small" type="primary">查看订单详细项目</el-button>
        </template>
      </el-table-column>

      <el-table-column label="订单备注" prop="remark" width="120"></el-table-column>

      <el-table-column label="退款原因"
                       prop="refund_reason_id"
                       width="120"
                       :filters="[{text: '买多了/不想买了', value: 3}, {text: '信息写错，重新下单', value: 4},
                       {text: '朋友/网上评价不好', value: 5}, {text: '计划有变，时间按排不上', value: 6},
                       {text: '其它', value: 7}]"
                       :filter-method="filterRefundReasonId"
                       filter-placement="bottom-end">
        <template slot-scope="scope">

          <el-tag
                  disable-transitions> {{scope.row.refund_reason_id|formatRefundReasonId}}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column label="退款原因备注" prop="refund_reason_remark" width="120"></el-table-column>



      <el-table-column label="订单跟进状态"
                       prop="handel_status"
                       width="120"
                       :filters="[{text: '待跟进', value: 0}, {text: '跟进中', value: 1}, {text: '跟进完毕', value: 2}]"
                       :filter-method="filterHandleStatus"
                       sortable
                       filter-placement="bottom-end">
        <template slot-scope="scope">

          <el-tag
                  :type="scope.row.handle_status === 0 ? 'warning' : scope.row.handle_status === 1 ? 'primary' : scope.row.handle_status === 2 ? 'success' : 'info'"
                  disable-transitions> {{scope.row.handle_status|formatHandleStatus}}
          </el-tag>
        </template>
      </el-table-column>

      <el-table-column label="订单跟进备注" prop="handle_remark" width="120"></el-table-column>


      <el-table-column label="创建时间" prop="create_time" width="120">
        <template slot-scope="scope">{{scope.row.create_time|formatDate}}</template>
      </el-table-column>

      <el-table-column label="更新时间" prop="update_time" width="120">
        <template slot-scope="scope">{{scope.row.update_time|formatDate}}</template>
      </el-table-column>

      <el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-button @click="notifyRefundOk(scope.row)" size="small" type="primary">通知已申请退款</el-button>
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

          <el-col :span="7">
            <el-form-item label="处理状态" prop="handle_status">
              <el-select v-model="formData.handle_status" placeholder="请选择处理状态" clearable :style="{width: '100%'}">
                <el-option v-for="(item, index) in handleStatusOptions" :key="index" :label="item.label"
                           :value="item.value" :disabled="item.disabled"></el-option>
              </el-select>
            </el-form-item>
          </el-col>

          <el-col :span="24">
            <el-form-item label="备注" prop="handle_remark">
              <el-input v-model="formData.handle_remark" type="textarea" placeholder="请输入处理备注" :maxlength="500"
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


    <el-dialog :before-close="closeRefundDialog" :visible.sync="dialogRefundFormVisible" title="退款短信通知弹窗操作">
      <el-row :gutter="15">
        <el-form ref="elForm" :model="refundFormData" :rules="refundRules" size="medium" label-width="100px">
          <el-col :span="7">
            <el-form-item label="订单id" prop="order_id">
              <el-input-number v-model="refundFormData.order_id" placeholder="订单id" :disabled='true'></el-input-number>
            </el-form-item>
          </el-col>

          <el-col :span="7">
            <el-form-item label="订单号" prop="out_trade_no">
              <el-input v-model="refundFormData.out_trade_no" placeholder="订单号" :disabled='true'></el-input>
            </el-form-item>
          </el-col>

          <el-col :span="18">
            <el-form-item label="套餐id" prop="pkg_id">
              <el-input-number v-model="refundFormData.pkg_id" placeholder="套餐id"></el-input-number>
            </el-form-item>
          </el-col>

          <el-col :span="18">
            <el-form-item label="套餐数量" prop="pkg_count">
              <el-input-number v-model="refundFormData.pkg_count" placeholder="套餐数量"></el-input-number>
            </el-form-item>
          </el-col>

          <el-col :span="18">
            <el-form-item label="退款金额" prop="amount">
              <el-input-number v-model="refundFormData.amount" placeholder="退款金额"></el-input-number>
            </el-form-item>
          </el-col>

        </el-form>
      </el-row>
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeRefundDialog">取 消</el-button>
        <el-button @click="enterRefundDialog" type="primary">确 定</el-button>
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
  import {formatTimeToStr} from "@/utils/data";
  import infoList from "@/components/mixins/infoList";
  import {formatPrice} from "@/utils/price";
  import {notifyRefundOk} from "@/api/notify"

  let that
  export default {
    name: "Order",
    mixins: [infoList],
    data() {
      return {
        listApi: getOrderList,
        dialogFormVisible: false,
        dialogRefundFormVisible: false,
        visible: false,
        type: "",
        formData: {
          id: null,
          status: null,
          refund_reason_remark: null,
          refund_reason_id: null,
          handle_status: null,
          handle_remark: null
        },
        refundFormData: {
          out_trade_no:null,pkg_id:null,pkg_count:1,amount:null,order_id:null,mobile:null
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
        refundRules: {
          pkg_id: [{
            required: true,
            message: '套餐id',
            trigger: 'blur'
          }],
          pkg_count: [{
            required: true,
            message: '套餐数量',
            trigger: 'blur'
          }],
          amount: [{
            required: true,
            message: '退款金额',
            trigger: 'blur'
          }],
        },
        statusOptions: [{
          "label": "待支付",
          "value": 0,
          "disabled": true
        }, {
          "label": "付款成功",
          "value": 2,
          "disabled": true
        }, {
          "label": "已退款",
          "value": 3
        }, {
          "label": "已过期/关闭",
          "value": 4,
          "disabled": true
        }, {
          "label": "体检完毕/待评价",
          "value": 5
        }],
        statusDict: new Map([[0, "待支付"], [2, "付款成功"], [3, "已退款"],
          [4, "已过期/关闭"], [5, "体检完毕/待评价"]]),
        handleStatusOptions: [
          {
            "label": "待跟进",
            "value": 0
          },
          {
            "label": "跟进中",
            "value": 1
          },
          {
            "label": "跟进完毕",
            "value": 2
          }
        ],
        handleStatusDict: new Map([[0, "待跟进"], [1, "跟进中"], [2, "跟进完毕"]]),
        refundReasonDict: new Map([[0, ""], [3, "买多了/不想买了"], [4, "信息写错，重新下单"], [5, "朋友/网上评价不好"], [6, "计划有变，时间按排不上"], [7, "其它"]])
      };
    },
    filters: {
      formatDate: function (time) {
        if (time != null && time != "") {
          var date = new Date(time * 1000);
          return formatTimeToStr(date, "yyyy-MM-dd hh:mm:ss");
        } else {
          return "";
        }
      },
      formatBoolean: function (bool) {
        if (bool != null) {
          return bool ? "是" : "否";
        } else {
          return "";
        }
      },
      formatOrderStatus: function (orderStatus) {
        return that.statusDict.get(orderStatus);
      },
      formatHandleStatus: function(handleStatus) {
        return that.handleStatusDict.get(handleStatus);
      },
      formatRefundReasonId: function(refundReasonId) {
        return that.refundReasonDict.get(refundReasonId)
      },
      formatPrice: formatPrice
    },
    methods: {
      //条件搜索前端看此方法
      onSubmit() {
        this.page = 1
        this.pageSize = 10
        // if (this.searchInfo.handle_status_filter !== 0 && this.searchInfo.handle_status_filter !== 1 && this.searchInfo.handle_status_filter !== 2) {
        //   this.searchInfo.handle_status_filter = undefined;
        // }
        if (this.searchInfo["handle_status_filter"] === "") {
          this.searchInfo.handle_status_filter = undefined;
        }
        this.getTableData()
      },
      ShowOrderItem(row) {
        this.$router.push({name: "orderItem", params: {order_id: row.id}, query: {out_trade_no: row.out_trade_no}})
      },
      clearFilter() {
        this.$refs.multipleTable.clearFilter();
      },
      filterStatus(value, row) {
        return row.status === value;
      },
      filterHandleStatus(value, row) {
        return row.handle_status === value
      },
      filterRefundReasonId(value, row) {
        return row.refund_reason_id === value
      },
      async notifyRefundOk(row) {
        this.refundFormData.out_trade_no = row.out_trade_no;
        this.refundFormData.order_id = row.id;
        this.refundFormData.mobile = row.mobile
        this.dialogRefundFormVisible = true;
      },
      async updateOrder(row) {
        const res = await findOrder({id: row.id});
        this.type = "update";
        if (res.code == 0) {
          this.formData = res.data.reorder;
          this.dialogFormVisible = true;
        }
      },
      closeDialog() {
        this.dialogFormVisible = false;
        this.formData = {

          id: null,
          out_trade_no: null,
          user_id: null,
          mobile: null,
          open_id: null,
          amount: null,
          refund_time: null,
          status: null,
          remark: null,
          handle_status: null,
          handle_remark: null,
          refund_reason_id: null,
          refund_reason_remark: null,
          create_time: null,
          update_time: null,
        };
      },
      closeRefundDialog() {
        this.dialogRefundFormVisible = false;
        this.refundFormData = {out_trade_no:null,pkg_id:null,pkg_count:null,amount:null,order_id:null,mobile:null}
      },
      async deleteOrder(row) {
        this.visible = false;
        const res = await deleteOrder({id: row.id});
        if (res.code == 0) {
          this.$message({
            type: "success",
            message: "删除成功"
          });
          await this.getTableData();
        }
      },
      async enterRefundDialog() {
        const res = await notifyRefundOk(this.refundFormData);
        if (res.code === 0) {
          this.$message({
            type: "success",
            message: "发送退款短信成功"
          })
          this.closeRefundDialog();
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
            type: "success",
            message: "创建/更改成功"
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

