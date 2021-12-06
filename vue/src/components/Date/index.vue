<template>
  <div class="filter-item">
    时间范围：
    <el-radio-group v-model="today">
      <el-radio-button label="0" >今天</el-radio-button>
      <el-radio-button label="1">昨天</el-radio-button>
      <el-radio-button label="6" >最近7天</el-radio-button>
      <el-radio-button label="29">最近30天</el-radio-button>
    </el-radio-group>
    <el-date-picker
      @change="searchData"
      v-model="date"
      type="daterange"
      align="right"
      unlink-panels
      range-separator="至"
      start-placeholder="开始日期"
      end-placeholder="结束日期"
      value-format="yyyy-MM-dd"
      :picker-options="pickerOptions">
    </el-date-picker>
  </div>
</template>

<script>
  import { dateFormat } from '@/utils/date'
  const day = 3600 * 1000 * 24
  export default {
    name: 'index',
    props:[
      'dates',
    ],
    data(){
      return {
        today:0,
        date:this.dates,
        pickerOptions: {
          shortcuts: [{
            text: '最近7天',
            onClick(picker) {
              const end = new Date();
              const start = new Date();
              start.setTime(start.getTime() - day * 7);
              picker.$emit('pick', [start, end]);
            }
          }, {
            text: '最近30天',
            onClick(picker) {
              const end = new Date();
              const start = new Date();
              start.setTime(start.getTime() - day * 30);
              picker.$emit('pick', [start, end]);
            }
          }, {
            text: '最近60天',
            onClick(picker) {
              const end = new Date();
              const start = new Date();
              start.setTime(start.getTime() - day * 60);
              picker.$emit('pick', [start, end]);
            }
          }, {
            text: '最近90天',
            onClick(picker) {
              const end = new Date();
              const start = new Date();
              start.setTime(start.getTime() - day * 90);
              picker.$emit('pick', [start, end]);
            }
          }, {
            text: '最近180天',
            onClick(picker) {
              const end = new Date();
              const start = new Date();
              start.setTime(start.getTime() - day* 180);
              picker.$emit('pick', [start, end]);
            }
          }, {
            text: '最近360天',
            onClick(picker) {
              const end = new Date();
              const start = new Date();
              start.setTime(start.getTime() - day * 360);
              picker.$emit('pick', [start, end]);
            }
          }]
        },
      }
    },
    watch:{
      today(newVal,oldVal) {
        this.setTime(newVal)
        this.$emit('changeDate',this.date)
        this.$emit('searchData',1)
      }
    },

    methods:{
      setTime(sum){
        const end = new Date();
        const start = new Date();
        start.setTime(start.getTime() - day * sum );
        this.date[0] = dateFormat("YYYY-mm-dd",start)
        this.date[1] = dateFormat("YYYY-mm-dd",end)
      },
      searchData(){
        this.$emit('changeDate',this.date)
        this.$emit('searchData',1)
      },
    },
    mounted() {
      this.setTime(1)
    }
  }
</script>

<style scoped>

</style>
