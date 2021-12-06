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

      <el-table-column align="center" label="类型" width="300">
        <template slot-scope="scope">
          {{scope.row.solut.name}}
        </template>
      </el-table-column>

      <el-table-column align="center" label="方案概述" width="300">
        <template slot-scope="scope">
          <img :src="scope.row.solut.fags" height="50px" @click="handlePictureCardPreview(scope.row.solut.fags)">
        </template>
      </el-table-column>

      <el-table-column align="center" label="方案价值" width="300">
        <template slot-scope="scope">
          <img :src="scope.row.solut.fajz" height="50px" @click="handlePictureCardPreview(scope.row.solut.fajz)">
        </template>
      </el-table-column>

      <el-table-column align="center" label="方案组成" >
        <template slot-scope="scope">
          <img :src="scope.row.solut.fazc" height="50px" @click="handlePictureCardPreview(scope.row.solut.fazc)">
        </template>
      </el-table-column>

      <el-table-column
        fixed="right"
        label="操作"
        width="200">
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
            @click.native.prevent="del(scope.row.solut.id)"
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
    <el-dialog :visible.sync="dialogVisible" @close="formClose" title="解决方案" style="position: absolute;z-index: 50">
      <el-form :model="form.solut" label-width="120px" label-position="left">

        <el-form-item label="类型">
          <el-input  v-model="form.solut.name"  placeholder="类型" />
        </el-form-item>

        <el-form-item label="方案概述">
          <upload-imgs v-if="dialogVisible" @handlePictureCardPreview="handlePictureCardPreview" @changeFileList="changefagsFileList" :data="fagsFileList" :limit="limit"></upload-imgs>
        </el-form-item>

        <el-form-item label="方案价值">
          <upload-imgs v-if="dialogVisible" @handlePictureCardPreview="handlePictureCardPreview" @changeFileList="changefajzFileList" :data="fajzFileList" :limit="limit"></upload-imgs>
        </el-form-item>

        <el-form-item label="方案组成">
          <upload-imgs v-if="dialogVisible" @handlePictureCardPreview="handlePictureCardPreview" @changeFileList="changefazcFileList" :data="fazcFileList" :limit="limit"></upload-imgs>
        </el-form-item>

      </el-form>

      <div v-for="(v,k,index)  in form.fatd" style="margin-top: 100px;box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);padding: 30px">

       <el-form :model="form.solut" label-width="150px" label-position="left" :label-suffix="createFormLableSuffix(k)">
          <el-form-item label="标题 ">
            <el-input    v-model="form.fatd[k].title" placeholder="标题" type="textarea" autosize />
          </el-form-item>
          <el-form-item label="图片 ">
            <upload-imgs v-if="dialogVisible" :index="k" @handlePictureCardPreview="handlePictureCardPreview" @changeFileList="changefatdFileList" :data="fatdFileList[k]" :limit="limit"></upload-imgs>
          </el-form-item>

          <el-form-item label="描述 ">
            <el-input    v-model="form.fatd[k].describe" placeholder="描述" type="textarea" autosize />
          </el-form-item>
         <el-form-item label="操作 ">
           <el-button icon="el-icon-minus" type="primary" class="filter-item" @click="removeFazcForm(k)"  >删除方案组成</el-button>
         </el-form-item>
        </el-form>
      </div>

      <div style="text-align:right;">
        <el-button type="danger" @click="formClose" icon="el-icon-close">取消</el-button>
        <el-button icon="el-icon-plus" type="primary" class="filter-item" @click="addFazcForm"  >增加方案组成</el-button>
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
  import {list,NameOptions,UpdateAction,Delete} from "../../api/solut"

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
          name:''
        },
        table:[],
        tableLoading:false,
        count:0,
        nameSelect:[],
        form:{
          solut:{
            id:0 ,
            name:'',
            fags:'',
            fajz:'',
            fazc:''
          },
          fatd:[

          ],
        },
        dialogVisible:false,
        fagsFileList: [],
        fajzFileList: [],
        fazcFileList: [],
        fatdFileList: [],
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
      removeFazcForm(i){
        this.form.fatd.splice(i,1)
        this.fatdFileList.splice(i,1)
      },
      addFazcForm(){
        this.form.fatd.push({
          title:'',
          img:'',
          describe:'',
          solu_id:this.form.solut.id
        })
        this.fatdFileList.push([])
      },
      changefatdFileList(v,index){
        this.fatdFileList[index] = v
      },
      createFormLableSuffix(index){
        return "(方案特点:"+Number(index+1)+")"
      },
      selectBlur(e) {
        this.form.solut.name = e.target.value
      },
      changefagsFileList(v){
        console.log("v",v)
        this.fagsFileList = v
      },
      changefazcFileList(v){
        this.fazcFileList = v
      },
      changefajzFileList(v){
        this.fajzFileList = v
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

        this.form.solut.fazc =  getImageName(this.fazcFileList,0)


        this.form.solut.fajz =  getImageName(this.fajzFileList,0)
        this.form.solut.fags =  getImageName(this.fagsFileList,0)

        console.log(this.fagsFileList,getImageName(this.fagsFileList,0))

        for(let k in this.fatdFileList){
          this.form.fatd[k].img = getImageName(this.fatdFileList[k],0)
        }
        console.log(this.form)

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
        this.param = ''
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

            for(let k in this.table){
              if(this.table[k].fatd == null){
                this.table[k].fatd = []
              }
            }
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
          for(let i in this.form.fatd){
            let tmp = []
            tmp.push({"url":this.form.fatd[i].img})
            this.fatdFileList[i] = tmp
          }

          console.log(this.fatdFileList,this.form.fatd)

          if (row.solut.fazc !=""){
            this.fazcFileList.push({"url":row.solut.fazc})
          }

          if (row.solut.fajz !=""){
            this.fajzFileList.push({"url":row.solut.fajz})
          }
          if (row.solut.fags !=""){
            this.fagsFileList.push({"url":row.solut.fags})
          }
        }
        this.dialogVisible = true
      },
      formClose(){
        this.form.solut = deepClone({})
        this.fatdFileList = []
        this.form.fatd = []
        this.fazcFileList = []
        this.fajzFileList = []
        this.fagsFileList = []
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
