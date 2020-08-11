<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="套餐id">
          <el-input placeholder="搜索条件" v-model="searchInfo.id" clearable></el-input>
        </el-form-item>
        <el-form-item label="医院id">
          <el-input placeholder="搜索条件" v-model="searchInfo.hospital_id" clearable></el-input>
        </el-form-item>
        <el-form-item label="套餐名字">
          <el-input placeholder="搜索条件" v-model="searchInfo.name" clearable></el-input>
        </el-form-item>
        <el-form-item label="套餐目标人群">
          <el-select v-model="searchInfo.target" placeholder="请选择套餐目标人群搜索" clearable :style="{width: '100%'}">
            <el-option v-for="(item, index) in targetOptions" :key="index" :label="item.label"
                       :value="item.value" :disabled="item.disabled"></el-option>
          </el-select>
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

      <el-table-column label="套餐id" prop="id" width="80"></el-table-column>
      <el-table-column label="所属医院" prop="hospital_id" width="80">
        <template slot-scope="scope">
          <div>{{ scope.row.hospital_id | formatHospitalName }}</div>
        </template>
      </el-table-column>

      <el-table-column label="套餐名字" prop="name" width="100"></el-table-column>

      <el-table-column label="套餐目标人群" prop="target" width="120">
        <template slot-scope="scope">
          <div>{{ scope.row.target | formatTarget }}</div>
        </template>
      </el-table-column>

      <el-table-column label="目标疾病" prop="disease" width="120">
        <template slot-scope="scope">
          <div class="fl-left left-mg-xs">
            <el-select v-model="scope.row.disease_ids" multiple placeholder="请选择" filterable>
              <el-option
                      v-for="item in diseaseOptions"
                      :key="item.id"
                      :label="item.name"
                      :value="item.id">
              </el-option>
            </el-select>
            <el-button @click="putPkgDisease(scope.row)" type="primary">确 定</el-button>
          </div>
        </template>
      </el-table-column>

      <el-table-column label="门市价" prop="price_original" width="120">
        <template slot-scope="scope">
          <div>{{ scope.row.price_original | formatPrice }}</div>
        </template>
      </el-table-column>

      <el-table-column label="真实价格" prop="price_real" width="120" sortable>
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

      <el-table-column label="套餐类别" prop="ctg_ids" width="170">
        <template slot-scope="scope">
          <div class="fl-left left-mg-xs">
            <el-select v-model="scope.row.ctg_ids" multiple placeholder="请选择" filterable>
              <el-option
                      v-for="item in ctgOptions"
                      :key="item.id"
                      :label="item.name"
                      :value="item.id">
              </el-option>
            </el-select>
            <el-button @click="putPkgCategory(scope.row)" type="primary">确 定</el-button>
          </div>

        </template>
      </el-table-column>

      <el-table-column label="套餐属性" prop="" width="140">
        <template slot-scope="scope">
          <el-button @click="postPkgAttr(scope.row)" size="small" type="primary">变更套餐属性</el-button>
        </template>
      </el-table-column>


      <el-table-column label="创建时间" prop="create_time" width="120" sortable>
        <template slot-scope="scope">
          <div>{{ scope.row.create_time | formatDate }}</div>
        </template>
      </el-table-column>

      <el-table-column label="更新时间" prop="update_time" width="120" sortable>
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
      <el-row :gutter="15">
        <el-form ref="elForm" :model="formData" :rules="rules" size="large" label-width="100px" label-height="100px">
          <el-col :span="17">
            <el-form-item label="套餐所属医院" prop="hospital_id">
              <el-select v-model="formData.hospital_id" placeholder="请选择套餐所属医院" clearable
                         :style="{width: '100%'}">
                <el-option v-for="(item, index) in hospitalOptions" :key="index" :label="item.name"
                           :value="item.id" :disabled="item.disabled"></el-option>
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
          <el-col :span="24">
            <el-form-item size="large">
              <el-button @click="resetForm">重置</el-button>
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
    createPackage,
    deletePackage,
    updatePackage,
    findPackage,
    getPackageList,
    upatePkgCtgRelation,
    upatePkgDiseaseRelation
  } from "@/api/pkg";  //  此处请自行替换地址

  import { getHospitalList } from "@/api/hospital";

  import {
    getPkgCategoryList,
  } from "@/api/pkg_category";

  import {
    getDiseaseList
  } from "@/api/disease"
  import {formatPrice} from "@/utils/price";

  import {formatTimeToStr} from "@/utils/data";
  import infoList from "@/components/mixins/infoList";
  import {mapGetters} from "vuex";

  const path = process.env.VUE_APP_BASE_API
  let that
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
        ctgOptions: [],
        formData: {
          id: null,
          hospital_id: null,
          name: null,
          target: null,
          price_original: null,
          price_real: null,
          avatar_url: null,
          brief: null,
          comment: null,
          tips: null,
          create_time: null,
          update_time: null,
          is_deleted: null,
          ctg_ids: [],
          disease_ids: []
        },
        rules: {
          hospital_id: [{
            required: true,
            message: '请选择套餐所属医院',
            trigger: 'change'
          }],
          name: [{
            required: true,
            message: '请输入套餐名',
            trigger: 'blur'
          }],
          target: [{
            required: true,
            message: '请选择套餐目标人群',
            trigger: 'change'
          }],
          price_original: [{
            required: true,
            message: '门市价',
            trigger: 'blur'
          }],
          price_real: [{
            required: true,
            message: '真实价格',
            trigger: 'blur'
          }],
          brief: [],
          comment: [],
          tips: [],
        },
        hospitalOptions: [],
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
        diseaseOptions: [],
        targetDict: new Map([[0, "不限"], [1, "男士"], [2, "女-未婚"], [3, "女-已婚"]]),
        hospitalDict: new Map()
      };
    },
    filters: {
      formatDate: function (time) {
        if (time != null && time != "") {
          let date = new Date(time * 1000);
          return formatTimeToStr(date, "yyyy-MM-dd hh:mm:ss");
        } else {
          return "";
        }
      },
      formatTarget: function (targetCode) {
        return that.targetDict.get(targetCode);
      },
      // TODO 医院少的时候将就着用， 待优化
      formatHospitalName: function(hospitalId) {
        if (!that.hospitalDict.size) {
          that.getHospitalDict();
        }
        return that.hospitalDict.get(hospitalId);
      },
      formatBoolean: function (bool) {
        if (bool != null) {
          return bool ? "是" : "否";
        } else {
          return "";
        }
      },
      formatPrice: formatPrice
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
      resetForm() {
        this.$refs['elForm'].resetFields()
      },
      postPkgAttr(row) {
        const pkgId = row.id
        this.$router.push({name: "pkgAttr", params: {pkg_id: pkgId}, query: {pkg_name: row.name}})
      },
      putPkgCategory(row) {
        upatePkgCtgRelation({
          id: row.id,
          ctg_ids: row.ctg_ids
        })
      },
      putPkgDisease(row) {
        upatePkgDiseaseRelation({
          id: row.id,
          disease_ids: row.disease_ids
        })
      },
      getPkgCategoryList() {
        getPkgCategoryList({page:1, pageSize: 500}).then(res => {
          if (res.code == 0) {
            this.ctgOptions = res.data.list;
          }
        })
      },
      getPkgDiseaseList() {
        getDiseaseList({page:1, pageSize: 500}).then(res => {
          if (res.code == 0) {
            this.diseaseOptions = res.data.list;
          }
        })
      },
      getHospitalList() {
        getHospitalList({page:1, pageSize: 1000}).then(res => {
          if (res.code === 0) {
            this.hospitalOptions = res.data.list;
          }
        })
      },
      getHospitalDict() {
        for (let i = 0; i < this.hospitalOptions.length; i++) {
          this.hospitalDict.set(this.hospitalOptions[i].id, this.hospitalOptions[i].name);
        }
      },
      handleAvatarSuccess(res) {
        for (let item of this.tableData) {
          if (item.id === res.data.id) {
            item.avatar_url = res.data.avatar_url;
            break;
          }
        }
        this.$message({
          type: "success",
          message: "上传成功",
        });
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
          await this.getTableData();
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
      this.getPkgCategoryList();
      this.getPkgDiseaseList();
      this.getHospitalList();
    },
    beforeCreate() {
      that = this
    }
  };
</script>

<style>
</style>
