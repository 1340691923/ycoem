<template>
  <div class="app-container">
<el-card class="box-card">

    <div class="filter-container">

      <el-select class="filter-item" @change="searchData(1)" placeholder="产品名称" style="width: 300px"  v-model="input.name" clearable filterable>
        <el-option
          v-for="(v,k,index) in nameSelect"
          :key="index"
          :label="v"
          :value="v"
        />
      </el-select>

      <el-button icon="el-icon-search" type="primary" class="filter-item" circle @click="searchData(1)" />
      <el-button icon="el-icon-plus" type="primary" class="filter-item" @click="formOpenProduct()" circle ></el-button>
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
        width="50"
      >
        <template slot-scope="scope">
          {{ scope.$index+1 }}
        </template>
      </el-table-column>

      <el-table-column
        prop="id"
        label="id"
        align="center"
        width="100"
      />
      <el-table-column
        prop="name"
        label="产品名称"
        align="center"
        width="500"
      />

      <el-table-column align="center" label="产品介绍图" >
        <template slot-scope="scope">
          <img :src="scope.row.img" height="100px" @click="handlePictureCardPreview(scope.row.img)">
        </template>
      </el-table-column>

      <el-table-column
        fixed="right"
        label="操作"
        width="500"

      >
        <template slot-scope="scope">
          <el-button
            type="text"
            size="small"

            @click.native.prevent="formOpenNode(scope.row,'点播系统')"
          >
            点播系统
          </el-button>
          <el-button
            type="text"
            size="small"
            @click.native.prevent="formOpenNode(scope.row,'特色功能')"
          >
            特色功能
          </el-button>
          <el-button
            type="text"
            size="small"
            @click.native.prevent="formOpenNode(scope.row,'硬件产品')"
          >
            硬件产品
          </el-button>

          <el-button
            type="text"
            size="small"
            @click.native.prevent="formOpenNode(scope.row,'配套设备')"
          >
            配套设备
          </el-button>

          <el-button
            type="text"
            size="small"
            @click.native.prevent="del(scope.row.id)"
          >
            移除
          </el-button>
        </template>
      </el-table-column>
    </el-table>

    <div class="pagination-container">
      <el-pagination
        background
        :current-page="input.page"
        :page-size="input.limit"
        layout="total, sizes, prev, pager, next, jumper"
        :total="count"
        @current-change="searchData"
        @size-change="handleSizeChange"
      />
    </div>

    <el-dialog :visible.sync="dialogVisible" @close="formClose" :title="productNodeTitle" style="position: absolute;z-index: 50">
      <el-form :model="form" label-width="120px" label-position="left">
        <el-form-item label="介绍图">
          <upload-imgs v-if="dialogVisible" @handlePictureCardPreview="handlePictureCardPreview" @changeFileList="changeNodeList" :data="productNodeImgs" :limit="limit"></upload-imgs>
        </el-form-item>
      </el-form>
      <div style="text-align:right;">
        <el-button icon="el-icon-close" type="danger" @click="formClose">取消</el-button>
        <el-button icon="el-icon-check" type="primary" @click="add">确定</el-button>
      </div>
    </el-dialog>

    <el-dialog :visible.sync="dialogProductVisible" @close="formCloseProduct" title="产品" style="position: absolute;z-index: 50">
      <el-form  label-width="120px" label-position="left">

        <el-form-item label="产品名">
          <el-input v-model="formProduct.name" placeholder="产品名" type="textarea" autosize  />
        </el-form-item>

        <el-form-item label="产品介绍图">
          <upload-imgs v-if="dialogProductVisible" @handlePictureCardPreview="handlePictureCardPreview" @changeFileList="changeFileList" :data="fileList" :limit="limit"></upload-imgs>
        </el-form-item>
      </el-form>
      <div style="text-align:right;">
        <el-button icon="el-icon-close" type="danger" @click="formCloseProduct">取消</el-button>
        <el-button icon="el-icon-check" type="primary" @click="addProduct">确定</el-button>
      </div>
    </el-dialog>

    <el-dialog :visible.sync="dialogImgVisible" style="position: absolute;z-index: 100" >
      <img width="100%"
           :src="dialogImageUrl"
           alt="">
    </el-dialog>
</el-card></div>
</template>

<script>
  import { deepClone } from '@/utils'
  import { getImageName } from '@/utils/upload'
  import {list,NameOptions,UpdateAction,Delete,FindNodeById,AddProductNode} from "../../api/product"

  export default {
    name: 'index',
    components:{
      UploadImgs :()=>import("@/components/UploadImgs"),
    },

    data(){
      return {
        productNodeForm:{
          product_id:0,
          types:"",
          imgs:"",
          id:0
        },
        productNodeTitle:"",
        dialogProductVisible:false,
        dialogImgVisible:false,
        dialogImageUrl:'',
        input:{
          page:1,
          limit:10,
          name:''
        },
        table:[],
        tableLoading:false,
        count:0,
        nameSelect:[],
        formProduct:{
          img:"",
          name:""
        },
        form:{
          id:'',
          product_id:'',
          imgs:[],
          name:''
        },
        dialogVisible:false,
        fileList: [],
        param: '', // 表单要提交的参数
        limit:9,
        productNodeImgs:[]
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
      selectBlur(e) {
        this.form.type = e.target.value
      },
      changeFileList(v){
        this.fileList = v
      },
      changeNodeList(v){
        this.productNodeImgs = v
      },
      del(id){
        Delete({id:id}).then(res=>{
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
        this.form.imgs = []
        console.log(this.productNodeImgs)
        for(let k in this.productNodeImgs){
          this.form.imgs.push(getImageName(this.productNodeImgs,k))
        }
        AddProductNode(this.form).then(res=>{
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
        this.param = ''
        this.fileList = []
      },
      addProduct(){
        this.formProduct.img = getImageName(this.fileList,0)
        UpdateAction(this.formProduct).then(res=>{
          if (res.code == 0) {
            this.$message({
              type: 'success',
              message: res.msg
            })
            this.nameOpt()
            this.searchData(1)
            this.formCloseProduct()
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
            this.count = Number(res.data.count)
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
      formOpenNode(row,name){
        this.productNodeTitle = name
        FindNodeById({name:name,product_id:row.id}).then(res=>{
          if (res.code != 0) {
            this.$message({
              type: 'error',
              message: res.msg
            })
            return
          }
          this.form = res.data
          for(let img of this.form.imgs){
            this.productNodeImgs.push({"url":img})
          }
          this.form.name = name
          this.form.product_id = row.id

        }).catch(err=>{
          this.$message({
            type: 'error',
            message: '网络异常'
          })
          console.log(err)
        })

        this.formOpen()
      },
      formOpen(){
        this.dialogVisible = true
      },

      formOpenProduct(){
        this.dialogProductVisible = true
      },

      formCloseProduct(){
        this.formProduct = deepClone({})
        this.fileList = []
        this.dialogProductVisible = false
      },

      formClose(){
        this.form = deepClone({})
        this.productNodeImgs = []
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
