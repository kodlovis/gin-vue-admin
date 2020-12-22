<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item>
          <div>
            <el-button @click="openDialog" type="primary" size="mini" slot="reference">添加指标</el-button>
            <el-button icon="el-icon-confirm" size="mini" slot="reference" type="danger" @click="removeEvaluationKpi">清空指标</el-button>
          </div>
        </el-form-item>
      </el-form>
    </div>
    <el-table
      :data="KpiData"
      border
      ref="multipleTable"
      stripe
      style="width: 100%"
      tooltip-effect="dark"
      @selection-change="handleSelectionChange"
      white-space: pre-line
    >
    <el-table-column label="指标名称">
      <template slot-scope="scope">
        <p v-for="(item,index) in scope.row.Kpis"
        :key="index">{{item.Name}}<br/></p>
      </template>
    </el-table-column>

    <el-table-column label="指标说明" width="360" type="textarea">
      <template slot-scope="scope">
        <span v-for="(item,index) in scope.row.Kpis"
        :key="index">{{item.Description}}<br/></span>
      </template>
    </el-table-column>

    <el-table-column label="指标算法" width="360" type="textarea">
      <template slot-scope="scope">
        <span v-for="(item,index) in scope.row.Kpis"
        :key="index">{{item.Category}}<br/></span>
      </template>
    </el-table-column>
    
    <el-table-column label="指标分数" prop="KpiScore" width="120"></el-table-column>
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

    <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="添加指标" :append-to-body="true" style="width: 90%,marigin:right">
    
    <div>
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item>
          <div>
            <el-button icon="el-icon-confirm" size="mini" slot="reference" type="primary" @click="kpiDataEnter">批量添加</el-button>
          </div>
        </el-form-item>
      </el-form>

    </div>
    <el-table
      :data="KpiScoreData"
      @selection-change="handleSelectionChange"
      border
      ref="multipleTable"
      stripe
      style="width: 100%"
      tooltip-effect="dark"
    >
    <el-table-column type="selection" width="55"></el-table-column>

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
    <el-table-column label="指标分数">
      <template slot-scope="scope">
        <el-form>
          <el-form-item v-for="(item,index) in scope.row.EvaluationKpis"
             :key="index">
            <el-input v-model="item.KpiScore" v-if="!scope.row.EvaluationKpis" clearable placeholder="请输入"
             >{{item.KpiScore}}</el-input>
            <el-input v-else clearable placeholder="请输入" v-model="item.KpiScore"
             ></el-input>
          </el-form-item>
        </el-form>
      </template>
    </el-table-column>
    </el-table>
    </el-dialog>
  </div>
</template>

<script>
import {
    getKpiList,
    addKpiEvaluation,
    getKpiByIds,
    getKpiEvaluation,
    getKpiScoreByIds
} from "@/api/pas/kpi";  //  此处请自行替换地址
import {
    removeEvaluationKpi,
} from "@/api/pas/evaluation";
import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";
export default {
  name: "kpis",
  mixins: [infoList],
  props: {
    row: {
      default: function() {
        return {}
      },
      type: Object
    },
  },
  data() {
    return {
      listApi: getKpiList,
      dialogFormVisible: false,
      visible: false,
      type: "",
      deleteVisible: false,
      multipleSelection: [],KpiScoreData: {
            Name:"",
            Description:"",
            Status:"",
            Category:"",
            Tags: 1,
            EvaluationKpis:0,
      },
      KpiData:{
            Name:"",
            ID:"",
            Category:"",
            EvaluationKpis:"",
            KpiScore:"",
      },
      EvaluationKpis:{
            KpiScore:"",
      }
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
    async openDialog() {
        this.dialogFormVisible = true;
    },
    async removeEvaluationKpi() {
        const res = await removeEvaluationKpi({ID:Number(this.row.ID)})
        if (res.code == 0) {
          this.$message({
            type: 'success',
            message: '删除成功'
          })
          this.deleteVisible = false
          this.getKpiScoreByIds()
        }
      },
    nodeChange(){
          this.needConfirm = true
      },
    // 暴露给外层使用的切换拦截统一方法
    enterAndNext(){
      this.kpiDataEnter()
    },

    //条件搜索前端看此方法
    onSubmit() {
      this.page = 1
      this.pageSize = 10
      this.getKpiScoreByIds()
    },
    handleSelectionChange(val) {
      this.multipleSelection = val
    },
    async kpiDataEnter(){
        const ids = []
        const score = []
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
            for (let i = 0; i < item.EvaluationKpis.length; i++) {
              var data = Number(item.EvaluationKpis[i].KpiScore)
              score.push(data)
            }
          })
          const checkArr = await getKpiByIds({ ids })
          const res = await addKpiEvaluation({
            kpis: checkArr.data.list,
            ID: this.row.ID,
            KpiScore: score
          })
          if(res.code == 0){
              this.$message({ type: 'success', message: "批量添加成功" })
          }
        this.dialogFormVisible = false;
        const ref = await getKpiEvaluation({ID:Number(this.row.ID)})
        if (ref.code == 0) {
        this.KpiData = ref.data.list;
        }
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
    const life = await getKpiScoreByIds({ID:Number(this.row.ID)});
    const res = await getKpiEvaluation({ID:Number(this.row.ID)})
    if (life.code == 0) {
      this.KpiScoreData = life.data.list
      this.KpiData = res.data.list
    }
    
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
.el-dialog{
  width: 80%;
}
</style>