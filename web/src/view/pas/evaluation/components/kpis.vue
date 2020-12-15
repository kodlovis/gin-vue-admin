<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item>
          <el-popover placement="top" v-model="deleteVisible" width="0">
          <el-button icon="el-icon-confirm" size="mini" slot="reference" type="primary" @click="kpiDataEnter">批量添加</el-button>
          </el-popover>
        </el-form-item>
      </el-form>
    </div>
    <el-table
      :data="tableData"
      @selection-change="handleSelectionChange"
      border
      ref="multipleTable"
      stripe
      style="width: 100%"
      tooltip-effect="dark"
    >
    <el-table-column type="selection" width="55" ref="multipleSelection" @check="nodeChange" :props="menuDefaultProps"></el-table-column>

    
    <el-table-column label="指标名称" prop="Name" width="120"></el-table-column> 
    
    <el-table-column label="指标说明" prop="Description" width="360" type="textarea"></el-table-column> 
    
    <el-table-column label="指标算法" prop="Category" width="360" type="textarea"></el-table-column> 

    <el-table-column label="标签名称">
      <template slot-scope="scope">
        <span v-for="(item,index) in scope.row.Tags"
        :key="index">{{item.Name}}<br/></span>
      </template>
    </el-table-column>
    <el-table-column label="标签类型">
      <template slot-scope="scope">
        <span v-for="(item,index) in scope.row.Tags"
        :key="index">{{item.Category}}<br/></span>
      </template>
    </el-table-column>
    </el-table>

    <el-pagination
      :current-page="page"
      :page-size="pageSize"
      :page-sizes="[3, 5,10, 30, 50, 100]"
      :style="{float:'right',padding:'20px'}"
      :total="total"
      @current-change="handleCurrentChange"
      @size-change="handleSizeChange"
      layout="total, sizes, prev, pager, next, jumper"
    ></el-pagination>
  </div>
</template>

<script>
import {
    getKpiList,
    addKpiEvaluation,
    getKpiByIds,
} from "@/api/pas/kpi";  //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";
export default {
  name: "kpi",
  mixins: [infoList],
  props: {
    row: {
      default: function() {
        return {}
      },
      type: Object
    }
  },
  data() {
    return {
      listApi: getKpiList,
      dialogFormVisible: false,
      visible: false,
      type: "",
      deleteVisible: false,
      menuDefaultProps: {
        label: function(data){
          return data.meta.title
        }
      },
      multipleSelection: [],formData: {
            Name:"",
            Description:"",
            Status:"",
            Category:"",
            Tags: 1,
      },
      evaluationData:{
        ID:"",
        Name:"",
      },
      tagData:{
           Name:"",
           ID:"",
           Category:"",
           Parentid:"",
      }, 
      tagColumn:{
           Name:"",
      },
    };
  },
  filters: {
    formatDate: function(time) {
      if (time != null && time != "") {
        var date = new Date(time);
        return formatTimeToStr(date, "yyyy-MM-dd hh:mm:ss");
      } else {
        return "";
      }
    },
    formatBoolean: function(bool) {
      if (bool != null) {
        return bool ? "是" :"否";
      } else {
        return "";
      }
    }
  },
  methods: {
    nodeChange(){
          this.needConfirm = true
      },
    // 暴露给外层使用的切换拦截统一方法
    enterAndNext(){
      this.kpiDataEnter()
    },
    async kpiDataEnter(){
          const ids = []
        if(this.multipleSelection.length == 0){
          this.$message({
            type: 'warning',
            message: '请选择要批量的数据'
          })
          return
        }
        this.multipleSelection &&
          this.multipleSelection.map(item => {
            ids.push(item.ID)
          })
          const checkArr = await getKpiByIds( ids )
          const res = await addKpiEvaluation({ 
            kpis: checkArr.data.list,
            ID: this.row.ID
          })
          if(res.code == 0){
              this.$message({ type: 'success', message: "批量添加成功" })
          }
      },
      //条件搜索前端看此方法
      onSubmit() {
        this.page = 1
        this.pageSize = 10        
        this.getTableData()
      },
      handleSelectionChange(val) {
        this.multipleSelection = val
      },
      
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
          Name:"",
          Description:"",
          Status:"",
          Category:"",
          Tags: 1,
      };
    },
  },
  async created() {
    await this.getTableData();
    // const res1 = await getKpiEvaluation({ ID: this.row.ID })
    // const kpis = res1.data.reEvaluation
  }
}

</script>

<style>

.el-table .cell {
  white-space: pre-line;
}
.tableX .el-table--scrollable-x .el-table__body-wrapper {
   padding: 0 0 5px 0;
    margin: 0 0 5px 0;
    overflow-x: auto;
  }
</style>