<template>
  <div class="app-container">
<el-card class="box-card">

    <div class="filter-container">

      <el-select class="filter-item" @change="searchData(1)" placeholder="服务名称" style="width: 300px"  v-model="input.name" clearable filterable>
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
        label="服务名称"
        align="center"
        width="500"
      />

      <el-table-column
        prop="type"
        label="类型"
        align="center"

      />

      <el-table-column
        fixed="right"
        label="操作"
        width="300"
      >
        <template slot-scope="scope">

          <el-button
            type="text"
            size="small"
            icon="el-icon-delete"
            @click.native.prevent="del(scope.row.id,scope.row.name)"
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

    <el-dialog :visible.sync="dialogVisible" @close="formClose" title="服务" style="position: absolute;z-index: 50">
      <el-form :model="form" label-width="120px" label-position="left">
        <el-form-item label="名称">
          <el-input v-model="form.name" placeholder="名称" style="width: 250px;" />
        </el-form-item>
        <el-form-item label="名称">
          <el-select class="filter-item"  placeholder="类型" style="width: 300px"  v-model="form.type" clearable filterable>
            <el-option
              v-for="(v,k,index) in typesOption"
              :key="index"
              :label="v"
              :value="v"
            />
          </el-select>
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
  import {list,NameOptions,UpdateAction,Delete} from "../../api/service"
  const downloadType="文件下载"
  const answerType="解答"
  export default {
    name: 'index',

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

        form:{
          id:0,
          type:'',
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
    methods:{
      selectBlur(e) {
        this.form.type = e.target.value
      },
      del(id,name){
        Delete({id:id,name:name}).then(res=>{
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
        if(row){
          this.form = deepClone(row)
        }
        this.dialogVisible = true
      },
      formClose(){
        this.form = deepClone({})
        this.dialogVisible = false
      },
    }
  }
</script>


<style scoped>

</style>
