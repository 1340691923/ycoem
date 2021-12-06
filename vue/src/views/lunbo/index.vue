<template>
  <div class="app-container">
<el-card class="box-card">

    <div class="filter-container">
      <el-button icon="el-icon-search" type="primary" class="filter-item" circle @click="searchData(1)" />
      <el-button icon="el-icon-plus" type="primary" class="filter-item" @click="formOpen()" circle ></el-button>
    </div>
    <el-table
:header-cell-style="{background:'#eef1f6',color:'#606266'}"
v-loading="tableLoading"
      :data="table"
      style="width: 100%"
    >
      <el-table-column
        label="序号"
        align="center"
        fixed
        width="200"
      >
        <template slot-scope="scope">
          {{ scope.$index+1 }}
        </template>
      </el-table-column>

      <el-table-column
        prop="name"
        label="模块名称"
        align="center"
        width="500"
      />

      <el-table-column align="center" label="展示图">
        <template slot-scope="scope">
          <img :src="scope.row.img" height="100px" @click="handlePictureCardPreview(scope.row.img)">
        </template>
      </el-table-column>

      <el-table-column
        fixed="right"
        label="操作"
        width="200"
      >
        <template slot-scope="scope">
          <el-button
            type="text"
            size="small"
            icon="el-icon-edit"
            @click.native.prevent="formOpen(scope.row)"
          >
            修改
          </el-button>
          <el-button
            type="text"
            size="small"
            icon="el-icon-delete"
            @click.native.prevent="del(scope.row.name)"
          >
            移除
          </el-button>
        </template>
      </el-table-column>
    </el-table>

    <el-dialog :visible.sync="dialogVisible" @close="formClose" title="模块图" style="position: absolute;z-index: 50">
      <el-form :model="form" label-width="120px" label-position="left">
        <el-form-item label="模块">
          <el-select class="filter-item"  placeholder="模块" style="width: 300px"  :disabled="isedit"  v-model="form.name" clearable filterable>
            <el-option
              v-for="(v,k,index) in nameSelect"
              :key="index"
              :label="v"
              :value="v"
            />
          </el-select>
        </el-form-item>

        <el-form-item label="展示图">
          <upload-imgs v-if="dialogVisible" @handlePictureCardPreview="handlePictureCardPreview" @changeFileList="changeFileList" :data="fileList" :limit="limit"></upload-imgs>
          <!--<el-input v-model="form.logo" placeholder="案例介绍图" />-->
        </el-form-item>

      </el-form>
      <div style="text-align:right;">
        <el-button type="danger" @click="formClose" icon="el-icon-close">取消</el-button>
        <el-button type="primary" @click="add" icon="el-icon-check">确定</el-button>
      </div>
    </el-dialog>
    <el-dialog :visible.sync="dialogImgVisible" style="position: absolute;z-index: 100" >
      <img width="100%"
           :src="dialogImageUrl"
           alt="">
    </el-dialog>
</el-card> </div>
</template>

<script>
  import { deepClone } from '@/utils'
  import { getImageName } from '@/utils/upload'
  import {list,NameOptions,UpdateAction,Delete} from "../../api/lunbo"
  export default {
    name: 'index',
    components:{
      UploadImgs :()=>import("@/components/UploadImgs"),
    },

    data(){
      return {
        isedit:false,
        dialogImgVisible:false,
        dialogImageUrl:'',
        input:{
          page:1,
          limit:10,
          title:'',
          addr:'',
          type:''
        },
        table:[],
        tableLoading:false,
        count:0,
        nameSelect:[],
        form:{
          name:'',
          img:'',
        },
        dialogVisible:false,
        fileList: [],
        param: '', // 表单要提交的参数
        limit:1
      }
    },
    mounted(){
      this.nameOpt()
      this.searchData(1)
    },
    computed:{
      uploadUrl(){
        return process.env.VUE_APP_BASE_API+"/api/upload/img"
      }
    },
    methods:{

      changeFileList(v){
        this.fileList = v
      },
      del(name){
        Delete({name:name}).then(res=>{
          if (res.code == 0) {
            this.$message({
              type: 'success',
              message: res.msg
            })
            this.nameOpt()
            this.searchData(1)
          } else {
            this.$message({
              type: 'error',
              message: res.msg
            })
          }
        }).catch(err=>{
          this.$message({
            type: 'error',
            message: '网络异常'
          })
          console.log(err)
        })
      },
      add(){
        if (this.form.name == ""){
          this.$message({
            type: 'error',
            message: "模块名不能为空"
          })
          return
        }
        this.form.img = getImageName(this.fileList,0)
        UpdateAction(this.form).then(res=>{
          if (res.code == 0) {
            this.$message({
              type: 'success',
              message: res.msg
            })
            this.nameOpt()
            this.searchData(1)
            this.formClose()
          } else {
            this.$message({
              type: 'error',
              message: res.msg
            })
          }
        }).catch(err=>{
          this.$message({
            type: 'error',
            message: '网络异常'
          })
          console.log(err)
        })
        this.fileList = []
      },
      nameOpt(){
        NameOptions().then(res => {
          this.nameSelect = res.data.list
        }).catch(err => {
          this.$message({
            type: 'error',
            message: '网络异常'
          })
          console.log(err)
        })
      },
      searchData(page){
        this.tableLoading = true
        !page ? this.input.page = 1 : this.input.page = page
        list(this.input).then(res => {
          if (res.code == 0) {
            this.table = res.data.list
            this.tableLoading = false
          } else {
            this.$message({
              type: 'error',
              message: res.msg
            })
          }
          this.tableLoading = false
        }).catch(err => {
          this.tableLoading = false
          this.$message({
            type: 'error',
            message: '网络异常'
          })
          console.log(err)
        })
      },
      handleSizeChange(v) {
        this.input.limit = v
        this.searchData(1)
      },
      formOpen(row){
        if(row){
          this.isedit = true
          if (row.img!=""){
            this.fileList.push({"url":row.img})
          }
          this.form = row
        }else{
          this.isedit = false
        }

        this.dialogVisible = true
      },
      formClose(){
        this.form = deepClone({})
        this.fileList = []
        this.dialogVisible = false
      },
      handlePictureCardPreview (file) {
        if (file.url){
          this.dialogImageUrl = file.url
        }else{
          this.dialogImageUrl = file
        }
        this.dialogImgVisible = true
      },
    }
  }
</script>


<style scoped>

</style>
