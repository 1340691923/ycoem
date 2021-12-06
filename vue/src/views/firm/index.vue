<template>
  <div class="app-container">
<el-card class="box-card">

    <div class="filter-container">

      <el-select class="filter-item" @change="searchData(1)" placeholder="类型" style="width: 300px"  v-model="input.name" clearable filterable>
        <el-option
          v-for="(v,k,index) in nameSelect"
          :key="index"
          :label="v"
          :value="v"
        />
      </el-select>

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
        label="类型"
        align="center"
        width="100"
      />

      <el-table-column
        prop="title"
        label="标题"
        align="center"
        width="100"
      />

      <el-table-column align="center" label="公众号文章链接" width="300">
        <template slot-scope="scope">
          <a :href="scope.row.wx_link" target="_blank">{{scope.row.wx_link}}</a>
        </template>
      </el-table-column>

      <el-table-column align="center" label="内容" >
        <template slot-scope="scope">
             <el-popover
                      placement="top-start"
                      title="内容"
                      width="600"
                      trigger="hover"
                    >
                      <div>{{ scope.row.content }}</div>
                      <span slot="reference">{{
                        scope.row.content.substr(0, 50) + "..." + ''
                      }}</span>
                    </el-popover>
        </template>
      </el-table-column>


      <!--<el-table-column align="center" label="图片" width="300">
        <template slot-scope="scope">
          <img :src="scope.row.img" height="50px" @click="handlePictureCardPreview(scope.row.img)">
        </template>
      </el-table-column>-->
      <el-table-column
        fixed="right"
        label="操作"
        width="300"
      >
        <template slot-scope="scope">
          <el-button
            icon="el-icon-edit"
            type="text"
            size="small"
            @click.native.prevent="formOpen(scope.row)"
          >
            修改
          </el-button>
          <el-button
            type="text"
            icon="el-icon-delete"
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

    <el-dialog :visible.sync="dialogVisible" @close="formClose" title="关于公司" style="position: absolute;z-index: 50">
      <el-form :model="form" label-width="120px" label-position="left">
        <el-form-item label="类型">
          <el-select class="filter-item"  placeholder="类型" style="width: 300px" @blur="selectBlur"  v-model="form.name" clearable filterable>
            <el-option
              v-for="(v,k,index) in nameSelect"
              :key="index"
              :label="v"
              :value="v"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="标题">
          <el-input v-model="form.title" placeholder="标题" style="width: 250px;" />
        </el-form-item>
        <el-form-item label="公众号文章链接">
          <el-input v-model="form.wx_link" placeholder="公众号文章链接" type="textarea" autosize />
        </el-form-item>
        <el-form-item label="内容">
          <el-input v-model="form.content" type="textarea"
                    autosize placeholder="内容"  />
        </el-form-item>
        <el-form-item label="图片">
          <upload-imgs v-if="dialogVisible" @handlePictureCardPreview="handlePictureCardPreview" @changeFileList="changeFileList" :data="fileList" :limit="limit"></upload-imgs>
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
</el-card></div>
</template>

<script>
  import { deepClone } from '@/utils'
  import { getImageName } from '@/utils/upload'
  import {list,NameOptions,UpdateAction,Delete} from "../../api/firm"
  const downloadType="文件下载"
  const answerType="解答"
  const defaultForm = {
    id:0,
    title:'',
    img:[],
    wx_link:'',
    content:'',
    name:''
  }
  export default {
    name: 'index',
    components:{
      UploadImgs :()=>import("@/components/UploadImgs"),
    },
    data(){
      return {
        typesOption:[downloadType,answerType],
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
        form:defaultForm,
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
    methods:{
      selectBlur(e) {
        this.form.name = e.target.value
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
        this.form.img = []
        for(let k in this.fileList){
          this.form.img.push(getImageName(this.fileList,k))
        }

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
      formOpen(row){
        console.log("row",row)
        if(row){
          console.log(row.img.split(","))
          if (row.img !=""){
            for(let img of row.img.split(",")){
              this.fileList.push({"url":img})
            }
          }
         /* if (row.img!=""){
            this.fileList.push({"url":row.img})
          }*/
          this.form = row
        }
        console.log(this.nameSelect)
        this.dialogVisible = true
      },
      formClose(){
        this.fileList = []
        this.form = Object.assign({}, defaultForm)
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
      changeFileList(v){
        this.fileList = v
      },
    }
  }
</script>


<style scoped>

</style>
