<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="套餐id">
          <el-input placeholder="搜索条件" v-model="searchInfo.id"></el-input>
        </el-form-item>
        <el-form-item label="医院id">
          <el-input placeholder="搜索条件" v-model="searchInfo.hospital_id"></el-input>
        </el-form-item>
        <el-form-item label="套餐名字">
          <el-input placeholder="搜索条件" v-model="searchInfo.name"></el-input>
        </el-form-item>
        <el-form-item label="套餐目标人群">
          <el-input placeholder="搜索条件" v-model="searchInfo.target"></el-input>
        </el-form-item>
        <el-form-item label="目标疾病">
          <el-input placeholder="搜索条件" v-model="searchInfo.disease"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog" type="primary">新增体检套餐</el-button>
        </el-form-item>
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

      <el-table-column label="套餐id" prop="id" width="120"></el-table-column>
      <el-table-column label="医院id" prop="hospital_id" width="120"></el-table-column>

      <el-table-column label="套餐名字" prop="name" width="120"></el-table-column>

      <el-table-column label="套餐目标人群" prop="target" width="120"></el-table-column>

      <el-table-column label="目标疾病" prop="disease" width="120"></el-table-column>

      <el-table-column label="门市价" prop="price_original" width="120">
        <template slot-scope="scope">
          <div>{{ scope.row.price_original | formatPrice }}</div>
        </template>
      </el-table-column>

      <el-table-column label="真实价格" prop="price_real" width="120">
        <template slot-scope="scope">
          <div>{{ scope.row.price_real | formatPrice }}</div>
        </template>
      </el-table-column>

      <el-table-column label="套餐头像" prop="avatar_url" width="120">
        <template slot-scope="scope">
          <div class="fl-left left-mg-xs">
            <el-upload
                    :headers="{'x-token':token}"
                    :data="{'id': scope.row.id}"
                    :on-success="handleAvatarSuccess"
                    :show-file-list="false"
                    :action="`${path}/pkg/uploadAvatar`"
                    class="avatar-uploader"
                    name="avatar"
            >
              <img :src="scope.row.avatar_url" :alt="scope.row.alt" height="80" width="80" class="avatar"
                   v-if="scope.row.avatar_url"/>
              <i class="el-icon-plus avatar-uploader-icon" v-else></i>
            </el-upload>

            <!-- <el-avatar :size="120" :src="userInfo.headerImg" shape="square"></el-avatar> -->
          </div>

        </template>
      </el-table-column>

      <el-table-column label="套餐概述" prop="brief" width="120"></el-table-column>

      <el-table-column label="套餐备注" prop="comment" width="120"></el-table-column>

      <el-table-column label="套餐温馨提示" prop="tips" width="120"></el-table-column>

      <el-table-column label="套餐类别" prop="ctg_ids" width="120">
        <template slot-scope="scope">
          <div class="fl-left left-mg-xs">
            <el-select v-model="scope.row.ctg_ids" multiple placeholder="请选择" @blur="putPkgCategory(scope.row)">
              <el-option
                      v-for="item in ctgOptions"
                      :key="item.id"
                      :label="item.name"
                      :value="item.id">
              </el-option>
            </el-select>
          </div>

        </template>
      </el-table-column>

      <el-table-column label="创建时间" prop="create_time" width="120">
        <template slot-scope="scope">
          <div>{{ scope.row.create_time | formatDate }}</div>
        </template>
      </el-table-column>

      <el-table-column label="更新时间" prop="update_time" width="120">
        <template slot-scope="scope">
          <div>{{ scope.row.update_time | formatDate }}</div>
        </template>
      </el-table-column>

      <el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-button @click="updatePackage(scope.row)" size="small" type="primary">变更</el-button>
          <el-popover placement="top" width="160" v-model="scope.row.visible">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin: 0">
              <el-button size="mini" type="text" @click="scope.row.visible = false">取消</el-button>
              <el-button type="primary" size="mini" @click="deletePackage(scope.row)">确定</el-button>
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

      <el-form ref="elForm" :model="formData" :rules="rules" size="large" label-width="100px" label-height="100px">
        <el-col :span="17">
          <el-form-item label="套餐所属医院" prop="hospital_id">
            <el-select v-model="formData.hospital_id" placeholder="请选择套餐所属医院" clearable
                       :style="{width: '100%'}">
              <el-option v-for="(item, index) in hospital_idOptions" :key="index" :label="item.label"
                         :value="item.value" :disabled="item.disabled"></el-option>
            </el-select>
          </el-form-item>
        </el-col>
        <el-col :span="17">
          <el-form-item label="套餐名" prop="name">
            <el-input v-model="formData.name" placeholder="请输入套餐名" :maxlength="64" clearable
                      :style="{width: '100%'}"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="17">
          <el-form-item label="套餐目标人群" prop="target">
            <el-select v-model="formData.target" placeholder="请选择套餐目标人群" clearable :style="{width: '100%'}">
              <el-option v-for="(item, index) in targetOptions" :key="index" :label="item.label"
                         :value="item.value" :disabled="item.disabled"></el-option>
            </el-select>
          </el-form-item>
        </el-col>
        <el-col :span="17">
          <el-form-item label="检测目标高发疾病" prop="disease">
            <el-select v-model="formData.disease" placeholder="请选择检测目标高发疾病" clearable
                       :style="{width: '100%'}">
              <el-option v-for="(item, index) in diseaseOptions" :key="index" :label="item.label"
                         :value="item.value" :disabled="item.disabled"></el-option>
            </el-select>
          </el-form-item>
        </el-col>
        <el-col :span="17">
          <el-form-item label="门市价" prop="price_original">
            <el-input-number v-model="formData.price_original" placeholder="门市价"></el-input-number>
          </el-form-item>
        </el-col>
        <el-col :span="17">
          <el-form-item label="真实价格" prop="price_real">
            <el-input-number v-model="formData.price_real" placeholder="真实价格"></el-input-number>
          </el-form-item>
        </el-col>
        <el-col :span="17">
          <el-form-item label="套餐概述" prop="brief">
            <el-input v-model="formData.brief" type="textarea" placeholder="请输入套餐概述"
                      :autosize="{minRows: 4, maxRows: 4}" :style="{width: '100%'}"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="17">
          <el-form-item label="套餐备注" prop="comment">
            <el-input v-model="formData.comment" type="textarea" placeholder="请输入套餐备注"
                      :autosize="{minRows: 4, maxRows: 4}" :style="{width: '100%'}"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="17">
          <el-form-item label="套餐温馨提示" prop="tips">
            <el-input v-model="formData.tips" type="textarea" placeholder="请输入套餐温馨提示"
                      :autosize="{minRows: 4, maxRows: 4}" :style="{width: '100%'}"></el-input>
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
    createPackage,
    deletePackage,
    updatePackage,
    findPackage,
    getPackageList,
    upatePkgCtgRelation,
    getPkgAttrByPkgId
  } from "@/api/pkg";  //  此处请自行替换地址
  import {
    getPkgCategoryList,
  } from "@/api/pkg_category";

  import {formatTimeToStr} from "@/utils/data";
  import infoList from "@/components/mixins/infoList";
  import {mapGetters} from "vuex";

  const path = process.env.VUE_APP_BASE_API
  export default {
    name: "Package",
    mixins: [infoList],
    data() {
      return {
        listApi: getPackageList,
        dialogFormVisible: false,
        visible: false,
        type: "",
        path: path,
        innerTableData: [],
        ctgOptions: [],
        formData: {
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
          is_deleted: null,
          ctg_ids: []
        },
        hospital_idOptions: [{
          "label": "美年大健康",
          "value": 0
        }, {
          "label": "茂名市人民医院",
          "value": 0
        }],
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
        diseaseOptions: [{
          "label": "不限",
          "value": 0
        }, {
          "label": "食物不耐受检测",
          "value": 1
        }, {
          "label": "骨关节疾病体检",
          "value": 2
        }, {
          "label": "健康防癌体检",
          "value": 3
        }, {
          "label": "幽门螺旋杆菌检测",
          "value": 4
        }, {
          "label": "甲状腺检测",
          "value": 5
        }, {
          "label": "糖尿病检测",
          "value": 6
        }]
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
      formatPrice: function (price) {
        return "" + price / 100 + " 元"
      }
    },
    computed: {
      ...mapGetters('user', ['userInfo', 'token'])
    },
    methods: {
      //条件搜索前端看此方法
      onSubmit() {
        this.page = 1
        this.pageSize = 10
        this.getTableData()
      },
      putPkgCategory(row) {
        upatePkgCtgRelation({
          id: row.id,
          ctg_ids: row.ctg_ids
        })
      },
      getPkgCategoryList() {
        getPkgCategoryList({}).then(res => {
          if (res.code == 0) {
            this.ctgOptions = res.data.list;
          }
        })
      },
      handleAvatarSuccess(res) {
        for (let item of this.tableData) {
          if (item.id === res.data.id) {
            item.avatar_url = res.data.avatar_url;
            break;
          }
        }
      },
      async updatePackage(row) {
        const res = await findPackage({id: row.id});
        this.type = "update";
        if (res.code == 0) {
          this.formData = res.data.repkg;
          this.dialogFormVisible = true;
        }
      },
      closeDialog() {
        this.dialogFormVisible = false;
        this.formData = {
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
          is_deleted: null,
        };
      },
      async deletePackage(row) {
        this.visible = false;
        const res = await deletePackage({id: row.id});
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
            res = await createPackage(this.formData);
            break;
          case "update":
            res = await updatePackage(this.formData);
            break;
          default:
            res = await createPackage(this.formData);
            break;
        }
        if (res.code == 0) {
          this.$message({
            type: "success",
            message: "创建/更改成功"
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
      this.getPkgCategoryList();
    }
  };
</script>

<style>
</style>