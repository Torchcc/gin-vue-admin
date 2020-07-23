<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">    
        <el-form-item label="类别名称">
          <el-input placeholder="搜索条件" v-model="searchInfo.name"></el-input>
        </el-form-item>      
        <el-form-item label="id">
          <el-input placeholder="搜索条件" v-model="searchInfo.id"></el-input>
        </el-form-item>    
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog" type="primary">新增高发疾病项</el-button>
        </el-form-item>
        <el-form-item>
          <el-popover placement="top" v-model="deleteVisible" width="160">
            <p>确定要删除吗？</p>
              <div style="text-align: right; margin: 0">
                <el-button @click="deleteVisible = false" size="mini" type="text">取消</el-button>
                <el-button @click="onDelete" size="mini" type="primary">确定</el-button>
              </div>
            <el-button icon="el-icon-delete" size="mini" slot="reference" type="danger">批量删除</el-button>
          </el-popover>
        </el-form-item>
      </el-form>
    </div>
    <el-table
      :data="tableData"
      @selection-change="handleSelectionChange"
      border
      ref="multipleTable"
      stripe
      style="width: 100%"
      tooltip-effect="dark"
    >
    <el-table-column type="selection" width="55"></el-table-column>

    <el-table-column label="id" prop="id" width="120"></el-table-column>
    <el-table-column label="高发疾病名称" prop="name" width="120"></el-table-column>
    <el-table-column label="创建时间" prop="create_time" width="120">
      <template slot-scope="scope">{{scope.row.create_time|formatDate}}</template>
    </el-table-column>
    <el-table-column label="更新时间" prop="update_time" width="120">
      <template slot-scope="scope">{{scope.row.update_time|formatDate}}</template>
    </el-table-column>
    

      <el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-button @click="updateDisease(scope.row)" size="small" type="primary">变更</el-button>
          <el-popover placement="top" width="160" v-model="scope.row.visible">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin: 0">
              <el-button size="mini" type="text" @click="scope.row.visible = false">取消</el-button>
              <el-button type="primary" size="mini" @click="deleteDisease(scope.row)">确定</el-button>
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
      <el-form ref="elForm" :model="formData" size="medium" label-width="100px">
        <el-form-item label="高发疾病名称" prop="name" required>
          <el-input v-model="formData.name" placeholder="请输入高发疾病名称" :maxlength="30" clearable
                    :style="{width: '100%'}"></el-input>
        </el-form-item>
        <el-form-item size="large">
          <el-button @click="resetForm">重置</el-button>
        </el-form-item>
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
    createDisease,
    deleteDisease,
    deleteDiseaseByIds,
    updateDisease,
    findDisease,
    getDiseaseList
} from "@/api/disease";  //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/data";
import infoList from "@/components/mixins/infoList";

export default {
  name: "Disease",
  mixins: [infoList],
  data() {
    return {
      listApi: getDiseaseList,
      dialogFormVisible: false,
      visible: false,
      type: "",
      deleteVisible: false,
      multipleSelection: [],
      formData: {
        create_time:null,is_deleted:null,name:null,update_time:null,id:null,
      }
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
    }
  },
  methods: {
      //条件搜索前端看此方法
      onSubmit() {
        this.page = 1
        this.pageSize = 10         
        this.getTableData()
      },
      handleSelectionChange(val) {
        this.multipleSelection = val
      },
    resetForm() {
      this.$refs['elForm'].resetFields()
    },
      async onDelete() {
        const ids = []
        this.multipleSelection &&
          this.multipleSelection.map(item => {
            ids.push(item.id)
          })
        const res = await deleteDiseaseByIds({ ids })
        if (res.code == 0) {
          this.$message({
            type: 'success',
            message: '删除成功'
          })
          this.deleteVisible = false
          this.getTableData()
        }
      },
    async updateDisease(row) {
      const res = await findDisease({ id: row.id });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.redisease;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
        
          create_time:null,
          is_deleted:null,
          name:null,
          update_time:null,
          id:null,
      };
    },
    async deleteDisease(row) {
      this.visible = false;
      const res = await deleteDisease({ id: row.id });
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
          res = await createDisease(this.formData);
          break;
        case "update":
          res = await updateDisease(this.formData);
          break;
        default:
          res = await createDisease(this.formData);
          break;
      }
      if (res.code == 0) {
        this.$message({
          type:"success",
          message:"创建/更改成功"
        })
        this.closeDialog();
        this.getTableData();
      }
    },
    openDialog() {
      this.type = "create";
      this.dialogFormVisible = true;
    }
  },
  created() {
    this.getTableData();
  }
};
</script>

<style>
</style>