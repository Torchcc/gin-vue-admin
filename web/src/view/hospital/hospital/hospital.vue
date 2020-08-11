<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="医院ID">
          <el-input placeholder="搜索条件" v-model="searchInfo.id"></el-input>
        </el-form-item>

        <el-form-item label="医院名称">
          <el-input placeholder="搜索条件" v-model="searchInfo.name"></el-input>
        </el-form-item>

        <el-form-item label="医院级别">
          <el-select v-model="searchInfo.level" placeholder="请选择医院等级" clearable :style="{width: '100%'}">
            <el-option v-for="(item, index) in levelOptions" :key="index" :label="item.label"
                       :value="item.value" :disabled="item.disabled"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="省行政id">
          <el-input placeholder="搜索条件" v-model="searchInfo.province_id"></el-input>
        </el-form-item>    
        <el-form-item label="市行政id">
          <el-input placeholder="搜索条件" v-model="searchInfo.city_id"></el-input>
        </el-form-item>    
        <el-form-item label="区行政id">
          <el-input placeholder="搜索条件" v-model="searchInfo.county_id"></el-input>
        </el-form-item>    
        <el-form-item label="镇行政id">
          <el-input placeholder="搜索条件" v-model="searchInfo.town_id"></el-input>
        </el-form-item>              
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog" type="primary">新增医院</el-button>
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

    <el-table-column label="医院ID" prop="id" width="120"></el-table-column>
    <el-table-column label="医院名称" prop="name" width="120"></el-table-column>

    <el-table-column label="医院头像" prop="avatar_url" width="120">
      <template slot-scope="scope">
        <div class="fl-left left-mg-xs">
          <el-upload
                  :headers="{'x-token':token}"
                  :data="{'id': scope.row.id}"
                  :on-success="handleAvatarSuccess"
                  :show-file-list="false"
                  :action="`${path}/hospital/uploadAvatar`"
                  class="avatar-uploader"
                  name="avatar"
          >
            <img :src="scope.row.avatar_url" :alt="scope.row.alt" height="80" width="80" class="avatar"
                 v-if="scope.row.avatar_url"/>
            <i class="el-icon-plus avatar-uploader-icon" v-else></i>
          </el-upload>

        </div>

      </template>
    </el-table-column>

    <el-table-column label="医院级别" prop="level_name" width="120"></el-table-column>
    
    <el-table-column label="省名称" prop="province_name" width="120"></el-table-column>
    
    <el-table-column label="市名称" prop="city_name" width="120"></el-table-column>
    
    <el-table-column label="区名称" prop="county_name" width="120"></el-table-column>
    
    <el-table-column label="镇/街道名称" prop="town_name" width="120"></el-table-column>
    <el-table-column label="医院地址，街道门牌" prop="address" width="120"></el-table-column>

      <!--    <el-table-column label="经度" prop="longitude" width="120"></el-table-column> -->
<!--    -->
<!--    <el-table-column label="纬度" prop="latitude" width="120"></el-table-column> -->
    
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
          <el-button @click="updateHospital(scope.row)" size="small" type="primary">变更</el-button>
          <el-popover placement="top" width="160" v-model="scope.row.visible">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin: 0">
              <el-button size="mini" type="text" @click="scope.row.visible = false">取消</el-button>
              <el-button type="primary" size="mini" @click="deleteHospital(scope.row)">确定</el-button>
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

      <div>
        <el-row :gutter="15">
          <el-form ref="elForm" :model="formData" :rules="rules" size="medium" label-width="100px">
            <el-col :span="12">
              <el-form-item label="医院名称" prop="name">
                <el-input v-model="formData.name" placeholder="请输入医院名称" clearable :style="{width: '100%'}">
                </el-input>
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="医院级别" prop="level">
                <el-select v-model="formData.level" placeholder="请选择医院级别" filterable clearable
                           :style="{width: '100%'}">
                  <el-option v-for="(item, index) in levelOptions" :key="index" :label="item.label"
                             :value="item.value" :disabled="item.disabled"></el-option>
                </el-select>
              </el-form-item>
            </el-col>


            <el-col :span="12">
            <el-form-item label="医院所在省">
            <el-select v-model="formData.province_id" placeholder="请选择省" @change="getCities(formData.province_id)" filterable clearable :style="{width: '100%'}">
              <el-option
                      v-for="item in provinces"
                      :key="item.id"
                      :label="item.name"
                      :value="item.id">
              </el-option>
            </el-select>
            </el-form-item>

            </el-col>

            <el-col :span="12">
              <el-form-item label="医院所在城市">
              <el-select v-model="formData.city_id" placeholder="请选择城市" @change="getCounties(formData.city_id)" filterable clearable :style="{width: '100%'}">
                <el-option
                        v-for="item in cities"
                        :key="item.id"
                        :label="item.name"
                        :value="item.id">
                </el-option>
              </el-select>
              </el-form-item>
            </el-col>

            <el-col :span="12">
              <el-form-item label="医院所在区/镇">
              <el-select v-model="formData.county_id" placeholder="请选择区/县" @change="getTowns(formData.county_id)" filterable clearable :style="{width: '100%'}">
                <el-option
                        v-for="item in counties"
                        :key="item.id"
                        :label="item.name"
                        :value="item.id">
                </el-option>
              </el-select>
              </el-form-item>
            </el-col>

            <el-col :span="12">
              <el-form-item label="医院所在镇/街道">

              <el-select v-model="formData.town_id" placeholder="请选择镇/街道" filterable clearable :style="{width: '100%'}">
                <el-option
                        v-for="item in towns"
                        :key="item.id"
                        :label="item.name"
                        :value="item.id">
                </el-option>
              </el-select>
              </el-form-item>
            </el-col>

            <el-col :span="24">
              <el-form-item label="详细地址" prop="address">
                <el-input v-model="formData.address" placeholder="请输入详细地址" clearable :style="{width: '100%'}">
                </el-input>
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
    createHospital,
    deleteHospital,
    deleteHospitalByIds,
    updateHospital,
    findHospital,
    getHospitalList
} from "@/api/hospital";  //  此处请自行替换地址
import { listRegions } from "@/api/region";
import { formatTimeToStr } from "@/utils/data";
import infoList from "@/components/mixins/infoList";
import {mapGetters} from "vuex";

const imgPath = process.env.VUE_APP_BASE_API

export default {
  name: "Hospital",
  mixins: [infoList],
  data() {
    return {
      provinces: [],
      cities: [],
      counties: [],
      towns: [],
      listApi: getHospitalList,
      dialogFormVisible: false,
      visible: false,
      type: "",
      path: imgPath,
      deleteVisible: false,
      multipleSelection: [],
      formData: {
        name:null,address:null,avatar_url:null,level:null,province_id:null,city_id:null,county_id:null,town_id:null,longitude:null,latitude:null,create_time:null,update_time:null,is_deleted:null,
      },
      levelOptions:[
        {"label": "不限",
        "value": 0},
        {"label": "公立三甲",
          "value": 1},
        {"label": "公立医院",
          "value": 2},
        {"label": "民营医院",
          "value": 3},
        {"label": "专业机构",
          "value": 4},
      ],
      rules: {
        name: [{
          required: true,
          message: '请输入医院名称',
          trigger: 'blur'
        }],
        level: [{
          required: true,
          message: '请选择医院级别',
          trigger: 'change'
        }],
        address: [{
          required: true,
          message: '请输入医院门牌楼栋',
          trigger: 'blur'
        }],
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
      async getProvinces(id = 0) {
        const res = await listRegions({ parent_id: id });
        this.provinces = res.data;
        this.cities = [];
        this.counties = [];
        this.towns = [];
        this.formData.province_id = 0;
        this.formData.city_id = 0;
        this.formData.county_id = 0;
        this.formData.town_id = 0;
      },
      async getCities(id = 0) {
        const res = await listRegions({ parent_id: id });
        this.cities = res.data;
        this.counties = [];
        this.towns = [];
        this.formData.city_id = 0;
        this.formData.county_id = 0;
        this.formData.town_id = 0;
      },
      async getCounties(id = 0) {
        const res = await listRegions({ parent_id: id });
        this.counties = res.data;
        this.towns = [];
        this.formData.county_id = 0;
        this.formData.town_id = 0;
      },
      async getTowns(id = 0) {
        const res = await listRegions({ parent_id: id });
        this.towns = res.data;
        this.formData.town_id = 0;
      },
      handleSelectionChange(val) {
        this.multipleSelection = val
      },
      handleAvatarSuccess(res) {
        for (let item of this.tableData) {
          if (item.id === res.data.id) {
            item.avatar_url = res.data.avatar_url;
            break;
          }
        }
      },
      async onDelete() {
        const ids = []
        this.multipleSelection &&
          this.multipleSelection.map(item => {
            ids.push(item.ID)
          })
        const res = await deleteHospitalByIds({ ids })
        if (res.code == 0) {
          this.$message({
            type: 'success',
            message: '删除成功'
          })
          this.deleteVisible = false
          await this.getTableData()
        }
      },
    async updateHospital(row) {
      const res = await findHospital({ id: row.id });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.rehospital;
        const cityRes = await listRegions({ parent_id: this.formData.province_id});
        this.cities = cityRes.data;
        const countyRes = await listRegions({ parent_id: this.formData.city_id});
        this.counties = countyRes.data;
        const townRes = await listRegions({ parent_id: this.formData.county_id});
        this.towns = townRes.data;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
        
          name:null,
          address:null,
          avatar_url:null,
          level:null,
          province_id:null,
          city_id:null,
          county_id:null,
          town_id:null,
          longitude:null,
          latitude:null,
          create_time:null,
          update_time:null,
          is_deleted:null,
      };
    },
    async deleteHospital(row) {
      this.visible = false;
      const res = await deleteHospital({ id: row.id });
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
          res = await createHospital(this.formData);
          break;
        case "update":
          res = await updateHospital(this.formData);
          break;
        default:
          res = await createHospital(this.formData);
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
    },
    resetForm() {
      this.$refs['elForm'].resetFields()
    }
  },
  created() {
    this.getTableData();
    this.getProvinces(0);

  }
};
</script>

<style>
</style>