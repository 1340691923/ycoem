<template>
  <div>
    <div class="app-container">
<el-card class="box-card">
      <el-form :model="form" ref="form" label-width="100px" class="demo-dynamic">
        <el-form-item
          prop="phone"
          label="手机号码"
        >
          <el-input type="textarea"
                    autosize v-model="form.phone" style="width:300px"></el-input>
        </el-form-item>
        <el-form-item
          prop="qq"
          label="QQ号"
        >
          <el-input style="width:300px" type="textarea"
                    autosize v-model="form.qq"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="submitForm('form')" icon="el-icon-check">提交</el-button>
          <el-button @click="resetForm()" icon="el-icon-delete">重置</el-button>
        </el-form-item>
      </el-form>
</el-card></div>
  </div>
</template>

<script>
  import {setphoneQQ,getphoneQQ} from "../../api/seo"
  export default {
    name: 'index',
    data(){
      return {
        form:{
          phone:"",
          qq:"",
        },
      }
    },
    mounted(){
      this.getSeo()
    },
    methods:{
      getSeo(){
        getphoneQQ().then(res=>{
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
          console.log(res,"res")
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
        this.$refs[formName].validate((valid) => {
          if (valid) {
            setphoneQQ(this.form).then(res=>{
              if (res.code != 0){
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
              this.getSeo()
            })
              .catch(err=>{
                console.log(err)
                this.$message({
                  message: "网络异常",
                  type: 'error'
                })
              })
          } else {
            this.$message({
              message: '请检查表单',
              type: 'error'
            })
            return false;
          }
        });
      },
      resetForm() {
        this.form.phone = ""
        this.form.qq = ""
      }
    }
  }
</script>

<style scoped>

</style>
