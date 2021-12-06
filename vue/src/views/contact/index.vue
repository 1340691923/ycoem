<template>
    <div class="app-container">
<el-card class="box-card">

      <div class="filter-container">

        <el-input class="filter-item" placeholder="地区(模糊查询)" style="width: 300px" @input="searchData(1)"  v-model="input.area"></el-input>
        <el-input class="filter-item" placeholder="大区经理" style="width: 300px" @input="searchData(1)"  v-model="input.manager"></el-input>
        <el-select class="filter-item" @change="searchData(1)" placeholder="事业部" style="width: 300px"  v-model="input.name" clearable filterable>
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
          label="事业部名称"
          align="center"
          width="200"
        />
        <el-table-column
          prop="area"
          label="地区"
          align="center"
          width="300"
        />
        <el-table-column
          prop="manager"
          label="大区经理"
          align="center"
          width="200"
        />
        <el-table-column
          prop="support"
          label="技术支持"
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
      <el-dialog :visible.sync="dialogVisible" @close="formClose" title="事业部信息">
        <el-form :model="form" label-width="120px" label-position="left">
          <el-form-item label="事业部名字">
            <el-select class="filter-item"  placeholder="事业部名字" style="width: 300px"  @blur="selectBlur"  v-model="form.name" filterable>
              <el-option
                v-for="(v,k,index) in nameSelect"
                :key="index"
                :label="v"
                :value="v"
              />
            </el-select>
          </el-form-item>
          <el-form-item label="地区">
            <el-input   v-model="form.area" placeholder="地区" />
          </el-form-item>
          <el-form-item label="大区经理">
            <el-input    v-model="form.manager" placeholder="请填写姓名+联系方式" />
          </el-form-item>
          <el-form-item label="技术支持">
            <el-input  v-model="form.support" placeholder="请填写姓名+联系方式" />
          </el-form-item>
        </el-form>
        <div style="text-align:right;">
          <el-button icon="el-icon-close" type="danger" @click="formClose">取消</el-button>
          <el-button icon="el-icon-check" type="primary" @click="add">确定</el-button>
        </div>
      </el-dialog>
</el-card></div>
</template>

<script>
  import { deepClone } from '@/utils'
  import {list,NameOptions,UpdateAction,Delete} from "../../api/concat"

  const defaultForm = {
    id:0 ,
    area:'',
    manager:'',
    name:'',
    support:''
  }

  export default {
    name: 'index',
    data(){
      return {
        input:{
          page:1,
          limit:10,
          area:'',
          manager:'',
          name:''
        },
        table:[],
        tableLoading:false,
        count:0,
        nameSelect:[],
        form:Object.assign({}, defaultForm),
        dialogVisible:false
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
          this.form = row
        }
        this.dialogVisible = true
      },
      formClose(){
        this.form = Object.assign({}, defaultForm)
        this.dialogVisible = false
      }
    }
  }
</script>


<style scoped>

</style>
