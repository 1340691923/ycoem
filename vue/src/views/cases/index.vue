<template>
  <div class="app-container">
<el-card class="box-card">

    <div class="filter-container">

      <el-input class="filter-item" @input="searchData(1)" placeholder="标题" style="width: 300px" v-model="input.title"></el-input>
      <el-input class="filter-item" @input="searchData(1)"  placeholder="位置" style="width: 300px"  v-model="input.addr"></el-input>
      <el-select class="filter-item" @change="searchData(1)" placeholder="类型" style="width: 300px"  v-model="input.type" clearable filterable>
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
        prop="type"
        label="类型"
        align="center"
        width="200"
      />

      <el-table-column align="center" label="案例介绍图" width="300">
        <template slot-scope="scope">
          <img :src="scope.row.logo" height="50px" @click="handlePictureCardPreview(scope.row.logo)">
        </template>
      </el-table-column>

      <el-table-column align="center" label="描述" width="400">
        <template slot-scope="scope">
             <el-popover
                      placement="top"
                      title="描述"
                      width="400"
                      trigger="hover"
                    >
                      <div>{{ scope.row.desc }}</div>
                      <span slot="reference">{{
                        scope.row.desc.substr(0, 50) + "..." + ''
                      }}</span>
                    </el-popover>
        </template>
      </el-table-column>

      <el-table-column
        prop="addr"
        label="位置"
        align="center"
        width="200"
      />
      <el-table-column
        prop="create_time"
        label="创建时间"
        align="center"
        width="200"
      />
      <el-table-column
        prop="update_time"
        label="修改时间"
        align="center"

      />
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
    <el-dialog :visible.sync="dialogVisible" @close="formClose" title="案例" style="position: absolute;z-index: 50">
      <el-form :model="form" label-width="120px" label-position="left">

        <el-form-item label="类型">
          <el-select class="filter-item"  placeholder="类型" style="width: 300px"  @blur="selectBlur"  v-model="form.type" clearable filterable>
            <el-option
              v-for="(v,k,index) in nameSelect"
              :key="index"
              :label="v"
              :value="v"
            />
          </el-select>
        </el-form-item>

        <el-form-item label="标题">
          <el-input v-model="form.title" @input="debug()" placeholder="标题" />
        </el-form-item>

        <el-form-item label="案例介绍图">
          <upload-imgs v-if="dialogVisible" @handlePictureCardPreview="handlePictureCardPreview" @changeFileList="changeFileList" :data="fileList" :limit="limit"></upload-imgs>
          <!--<el-input v-model="form.logo" placeholder="案例介绍图" />-->
        </el-form-item>
        <el-form-item label="地址">
          <el-input v-model="form.addr" placeholder="地址" />
        </el-form-item>

        <el-form-item label="描述">
          <el-input v-model="form.desc" placeholder="描述" type="textarea" autosize  />
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
  import {list,NameOptions,UpdateAction,Delete} from "../../api/cases"

  const defaultForm = {
    id:0 ,
    title:'',
    logo:'',
    name:'',
    desc:'',
    type:'',
    addr:''
  }

  export default {
    name: 'index',
    components:{
      UploadImgs :()=>import("@/components/UploadImgs"),
    },

    data(){
      return {
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
        form:Object.assign({}, defaultForm),
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
      selectBlur(e) {
        this.form.type = e.target.value
      },
      changeFileList(v){
        this.fileList = v
      },
      debug(){
        console.log(this.form)
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
        this.form.logo = getImageName(this.fileList,0)
        UpdateAction(this.form).then(res=>{
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
        this.formClose()
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
        if(row){
          if (row.logo!=""){
            this.fileList.push({"url":row.logo})
          }
          this.form = row
        }
        this.dialogVisible = true
      },
      formClose(){
        this.form = Object.assign({}, defaultForm)
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
