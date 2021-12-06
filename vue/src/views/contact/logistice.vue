<template>
  <div>
    <div class="app-container">
<el-card class="box-card">
      <el-form :model="form" ref="form" label-width="100px" class="demo-dynamic">
        <el-form-item
          prop="email"
          label="物流信息"
        >
          <el-input style="width:500px" v-model="form.concat" placeholder="请填写联系人名字和电话号码"></el-input>
        </el-form-item>

        <el-form-item
          prop="email"
          label="物流地址"
        >
          <el-input style="width:500px" v-model="form.area"></el-input>
        </el-form-item>

        <el-form-item>
          <el-button type="primary" @click="submitForm('form')" icon="el-icon-check">提交</el-button>

          <el-button @click="resetForm('form')" icon="el-icon-delete">重置</el-button>
        </el-form-item>
      </el-form>
</el-card></div>
  </div>
</template>

<script>
  import {GetLogisticeAction,SetLogisticeAction} from "../../api/concat"
  export default {
    name: 'index',
    data(){
      return {
        form:{
          area:'',
          concat:''
        }
      }
    },
    mounted(){
      this.Get()
    },
    methods:{
      Get(){
        GetLogisticeAction().then(res=>{
          if (res.code!=0){
            this.$message({
              message: res.msg,
              type: 'error'
            })
            return
          }
          this.$message({
            message: res.msg,
            type: 'success'
          })
          this.form = res.data
        })
          .catch(err=>{
            console.log(err)
            this.$message({
              message: "网络异常",
              type: 'error'
            })
          })
      },
      submitForm(formName) {
        SetLogisticeAction(this.form).then(res=>{
          if (res.code!=0){
            this.$message({
              message: res.msg,
              type: 'error'
            })
            return
          }
          this.$message({
            message: res.msg,
            type: 'success'
          })
          this.Get()
        })
          .catch(err=>{
            console.log(err)
            this.$message({
              message: "网络异常",
              type: 'error'
            })
          })
      },
      resetForm(formName) {
        this.form.area = ""
        this.form.concat = ""
      }
    }
  }
</script>

<style scoped>

</style>
