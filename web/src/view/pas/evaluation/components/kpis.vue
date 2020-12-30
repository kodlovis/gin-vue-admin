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

    <!-- <el-table-column label="评分人" width="120">
      <template slot-scope="scope">
        <span v-for="(item,index) in scope.row.Users"
        :key="index">{{item.nickName}}
        </span>
      </template>
    </el-table-column> -->
    
    <el-table-column label="评分人" width="120">
      <template slot-scope="scope">
          <el-cascader
            @change="(val)=>{handleOptionChange(val,scope.row)}"
            :options="userOptions"
            :rules="rules"
            clearable
            :props="{ checkStrictly: true,label:'nickName',value:'id',multiple: true,}"
            filterable
          ></el-cascader>
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

    <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="添加指标" :append-to-body="true" style="width: 90%,marigin:right">
    <el-table
      :data="kpiList"
      @selection-change="handleSelectionChange"
      border
      ref="multipleTable"
      stripe
      style="width: 100%"
      tooltip-effect="dark"
    >
    <el-table-column label="指标名称" prop="Name" width="120"></el-table-column> 
    
    <el-table-column label="指标说明" prop="Description" width="360" type="textarea"></el-table-column> 
    
    <!-- <el-table-column label="指标状态" prop="Status" width="120"></el-table-column>  -->
    
    <el-table-column label="指标算法" prop="Category" width="360" type="textarea"></el-table-column> 

     <el-table-column label="指标分数">
       
      <template slot-scope="scope">
          <el-input v-model="scope.row.EvaluationKpis.KpiScore" clearable placeholder="请输入"></el-input>
      </template>
    </el-table-column>
      <el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-button @click="kpiDataEnter(scope.row)" type="primary" size="mini" slot="reference" label="添加">添加</el-button>
        </template>
      </el-table-column>
    </el-table>
    </el-dialog>
  </div>
</template>

<script>
import {
    getKpiList,
    getKpiEvaluation,
} from "@/api/pas/kpi";  //  此处请自行替换地址
import {
    removeEvaluationKpi,
} from "@/api/pas/evaluation";
import {
    setUserEvaluation,
    createEvaluationKpi,
    deleteEvaluationKpi
} from "@/api/pas/evaluationKpi";
import {
    getUserList,
    getUserByIds
} from "@/api/user";
import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";
export default {
  name: "kpis",
  mixins: [infoList],
  ids:[],
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
      multipleOption: [],
      multipleSelection: [],
      KpiData: [],
      kpiList:{
            Name:"",
            ID:"",
            Category:"",
            KpiScore:"",
            EvaluationKpis:{
              KpiScore:"",
              evaluation_id:"",
              kpi_id:"",
        },
      },
      rules: {
        id: [
          { required: true, message: "请选择用户角色", trigger: "blur" }
        ]
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
          this.refreshEvalutationKpi()
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
    handleOptionChange(val,row) {
      const arr = this.$steamrollArray(val)
      console.log(arr,row.ID)
      this.multipleOption = val
      const ID = row.ID
      this.$message(ID+"123")
      this.changeUser(ID)
    },
    async kpiDataEnter(row){
        const res = await createEvaluationKpi({...this.row.EvaluationKpis,
        KpiScore: Number(row.EvaluationKpis.KpiScore),
        EvaluationId: this.row.ID,
        KpiId: row.ID,
          })
          if(res.code == 0){
              this.$message({ type: 'success', message: "添加成功" })
          }
          this.refreshEvalutationKpi()
      },
    async refreshEvalutationKpi(){
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
    
    setOptions(userData) {
      this.userOptions = [];
      this.ids = [];
      this.setUserOptions(userData, this.userOptions ,this.ids);
    },
    setUserOptions(UserData, optionsData ,ids) {
      UserData &&
        UserData.map(item => {
          // if (item.children && item.children.length) {
          //   const option = {
          //     ID: item.ID,
          //     nickName: item.nickName,
          //     children: []
          //   };
          //   this.setAuthorityOptions(item.children, option.children);
          //   optionsData.push(option);
          // } else {
            const option = {
              id: item.ID,
              nickName: item.nickName
            };
            optionsData.push(option);
            const idOption = {
              id: item.ID,
            };
            ids.push(idOption)
          //}
        });
    },
    async changeUser() {
      const ids = [];
      this.multipleOption &&
          this.multipleOption.map(item => {
            ids.push(Number(item))
          })
      const checkArr = await getUserByIds({ ids })
      const res = await setUserEvaluation({
        id: this.KpiData.ID,
        users: checkArr.data.list
      });
      if (res.code == 0) {
        this.$message({ type: "success", message: "角色设置成功" });
        this.refreshEvalutationKpi()
      }
    }
  },
  async created() {
    //const life = await getKpiScoreByIds({ID:Number(this.row.ID)});
    const life = await getKpiList({ID:Number(this.row.ID)});
    const res = await getKpiEvaluation({ID:Number(this.row.ID)});
      this.KpiData = res.data.list;
      this.kpiList = life.data.list;
    const user = await getUserList({ page: 1, pageSize: 999 });
      this.setOptions(user.data.list);
      this.refreshEvalutationKpi()
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