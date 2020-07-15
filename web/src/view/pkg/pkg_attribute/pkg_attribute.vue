<template>
  <div>
    <div><h1 class="pkg-name">{{$route.query.pkg_name}}</h1></div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item>
          <el-button @click="openDialog" type="primary">新增套餐属性</el-button>
        </el-form-item>
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

      <el-table-column label="属性id" prop="id" width="90"></el-table-column>

<!--      <el-table-column label="套餐id" prop="pkg_id" width="90"></el-table-column>-->

      <el-table-column label="属性类型"
                       prop="attr_type"
                       width="180"
                       :filters="[{text: '套餐项目', value: 1}, {text: '套餐需知', value: 2}, {text: '套餐流程', value: 3}]"
                       :filter-method="filterAttr"
                       sortable
                       filter-placement="bottom-end">
        <template slot-scope="scope">

          <el-tag
                  :type="scope.row.attr_type === 1 ? 'primary' : scope.row.attr_type === 2 ? 'success' : 'warning'"
                  disable-transitions> {{scope.row.attr_type|formatAttr}}
          </el-tag>
        </template>
      </el-table-column>

      <el-table-column label="排序号码" prop="order_no" sortable width="120"></el-table-column>

      <el-table-column label="属性名称" prop="name" width="180"></el-table-column>

      <el-table-column label="简介" prop="brief" width="180"></el-table-column>

      <el-table-column label="详细描述" prop="comment" width="600"></el-table-column>

      <el-table-column label="创建时间" prop="create_time" width="120">
        <template slot-scope="scope">{{scope.row.create_time|formatDate}}</template>
      </el-table-column>

      <el-table-column label="更新时间" prop="update_time" width="120">
        <template slot-scope="scope">{{scope.row.update_time|formatDate}}</template>
      </el-table-column>


      <el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-button @click="updatePkgAttr(scope.row)" size="small" type="primary">变更</el-button>
          <el-popover placement="top" width="160" v-model="scope.row.visible">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin: 0">
              <el-button size="mini" type="text" @click="scope.row.visible = false">取消</el-button>
              <el-button type="primary" size="mini" @click="deletePkgAttr(scope.row)">确定</el-button>
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
      <!--      一下div内容由表单生成器生成-->
      <div>
        <el-row :gutter="15">
          <el-form ref="elForm" :model="formData" :rules="rules" size="medium" label-width="100px">
            <el-col :span="13">
              <el-form-item label="套餐id" prop="pkg_id">
                <el-input-number v-model="formData.pkg_id" placeholder="套餐id" :disabled='true'></el-input-number>
              </el-form-item>
            </el-col>
            <el-col :span="13">
              <el-form-item label="属性类型" prop="attr_type">
                <el-select v-model="formData.attr_type" placeholder="请选择属性类型如,项目,需知,流程" clearable
                           :style="{width: '100%'}">
                  <el-option v-for="(item, index) in attrOptions" :key="index" :label="item.label"
                             :value="item.value" :disabled="item.disabled"></el-option>
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span="11">
              <el-form-item label="排序号码" prop="order_no">
                <el-input-number v-model="formData.order_no" placeholder="排序号码,小的排前面"></el-input-number>
              </el-form-item>
            </el-col>
            <el-col :span="24">
              <el-form-item label="属性名称" prop="name">
                <el-input v-model="formData.name" placeholder="请输入属性名称" :maxlength="30" clearable
                          :style="{width: '100%'}"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="24">
              <el-form-item label="属性简介" prop="brief">
                <el-input v-model="formData.brief" type="textarea" placeholder="请输入属性简介" :maxlength="500"
                          :autosize="{minRows: 6, maxRows: 6}" :style="{width: '100%'}"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="24">
              <el-form-item label="属性详细描述" prop="comment">
                <el-input v-model="formData.comment" type="textarea" placeholder="请输入属性详细描述" :maxlength="1000"
                          :autosize="{minRows: 10, maxRows: 10}" :style="{width: '100%'}"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="24">
              <el-form-item size="large">
                <el-button @click="resetForm">重置</el-button>
              </el-form-item>
            </el-col>
          </el-form>
        </el-row>
      </div>
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button @click="enterDialog" type="primary">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
  import {
    createPkgAttr,
    deletePkgAttr,
    updatePkgAttr,
    findPkgAttr,
    getPkgAttrList
  } from "@/api/pkg_attribute";  //  此处请自行替换地址
  import {formatTimeToStr} from "@/utils/data";
  import infoList from "@/components/mixins/infoList";
  let that;
  export default {
    name: "PkgAttr",
    mixins: [infoList],
    data() {
      return {
        attrDict: new Map([[1, "套餐项目"], [2, "套餐需知"], [3, "套餐流程"]]),
        listApi: getPkgAttrList,
        dialogFormVisible: false,
        attrOptions: [{"label": "套餐项目", "value": 1}, {"label": "套餐需知", "value": 2}, {"label": "套餐流程", "value": 3}],
        visible: false,
        type: "",
        formData: {
          pkg_id: this.$route.params.pkg_id,
          attr_type: '',
          order_no: 1,
          name: null,
          brief: null,
          comment: null,
        },
        rules: {
          pkg_id: [{
            required: true,
            message: '套餐id',
            trigger: 'blur'
          }],
          attr_type: [{
            required: true,
            message: '请选择属性类型如,项目,需知,流程',
            trigger: 'change'
          }],
          order_no: [{
            required: true,
            message: '排序号码,小的排前面',
            trigger: 'blur'
          }],
          name: [{
            required: true,
            message: '请输入属性名称',
            trigger: 'blur'
          }],
          brief: [{
            message: '请输入属性简介',
            trigger: 'blur'
          }],
          comment: [{
            message: '请输入属性详细描述',
            trigger: 'blur'
          }],
        },
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
      formatAttr: function (attr) {
        return that.attrDict.get(attr);
      },
      formatBoolean: function (bool) {
        if (bool != null) {
          return bool ? "是" : "否";
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
        this.$refs['elForm'].resetFields()
      },
      async getAllAttrByPkgId(page = this.page, pageSize = this.pageSize) {
        const pkg_id = this.$route.params.pkg_id
        const table = await this.listApi({page, pageSize, pkg_id})
        this.tableData = table.data.list
        this.total = table.data.total
        this.page = table.data.page
        this.pageSize = table.data.pageSize
      },
      clearFilter() {
        this.$refs.multipleTable.clearFilter();
      },
      filterAttr(value, row) {
        return row.attr_type === value;
      },
      async updatePkgAttr(row) {
        const res = await findPkgAttr({id: row.id});
        this.type = "update";
        if (res.code == 0) {
          this.formData = res.data.repkgAttr;
          this.dialogFormVisible = true;
        }
      },
      closeDialog() {
        this.dialogFormVisible = false;
        this.formData = {

          id: null,
          pkg_id: this.$route.params.pkg_id,
          attr_type: null,
          order_no: null,
          name: null,
          brief: null,
          comment: null,
          create_time: null,
          update_time: null,
          is_deleted: null,
        };
      },
      async deletePkgAttr(row) {
        this.visible = false;
        const res = await deletePkgAttr({id: row.id});
        if (res.code == 0) {
          this.$message({
            type: "success",
            message: "删除成功"
          });
          this.getTableData();
        }
      },
      async enterDialog() {
        let res;
        switch (this.type) {
          case "create":
            res = await createPkgAttr(this.formData);
            break;
          case "update":
            res = await updatePkgAttr(this.formData);
            break;
          default:
            res = await createPkgAttr(this.formData);
            break;
        }
        if (res.code == 0) {
          this.$message({
            type: "success",
            message: "创建/更改成功"
          })
          this.closeDialog();
          await this.getAllAttrByPkgId(1, 1000);
        }
      },
      openDialog() {
        this.type = "create";
        this.dialogFormVisible = true;
      }
    },
    created() {
      this.getAllAttrByPkgId(1, 1000);
    },
    beforeCreate() {
      that = this;
    }
  };
</script>

<style>
  .pkg-name {
    color: #5e6d82;
    -webkit-text-stroke: 1px black;
    letter-spacing: 0.04em;
    font-size: xx-large;
    text-align: center;
    margin: auto;
  }

</style>