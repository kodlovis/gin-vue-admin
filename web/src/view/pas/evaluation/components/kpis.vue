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

    <el-table-column label="评分人" width="120">
      <template slot-scope="scope">
        <span v-for="(item,index) in scope.row.Users"
        :key="index">{{item.nickName}}
        </span>
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
      :data="kpiList"
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
    
    <!-- <el-table-column label="指标状态" prop="Status" width="120"></el-table-column>  -->
    
    <el-table-column label="指标算法" prop="Category" width="360" type="textarea"></el-table-column> 

     <el-table-column label="指标分数">
          <el-input v-model="EvaluationKpiData.KpiScore" clearable placeholder="请输入"></el-input>
    </el-table-column>
    <el-table-column label="评分人" width="120">
      <template slot-scope="scope">
          <el-cascader
            @change="handleOptionChange"
            v-model="scope.id"
            :options="userOptions"
            :rules="rules"
            clearable
            :props="{ checkStrictly: true,label:'nickName',value:'id',multiple: true,}"
            filterable
            ref="cascaderAddr"
          ></el-cascader>
      </template>
    </el-table-column>
    <el-table-column>
      <el-button @click="kpiDataEnter" type="primary" size="mini" slot="reference">添加</el-button>
    </el-table-column>
    </el-table>
    </el-dialog>
  </div>
</template>

<script>
import {
    getKpiList,
    assignedKpiEvaluation,
    getKpiByIds,
    getKpiEvaluation,
} from "@/api/pas/kpi";  //  此处请自行替换地址
import {
    removeEvaluationKpi,
} from "@/api/pas/evaluation";
import {
    getEvaluationKpiList,
    setUserEvaluation
} from "@/api/pas/evaluationKpi";
import {
    getUserList
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
            Users:{
              ID:[]
            }
      },
      EvaluationKpiData:{
        ids:["1","2"]
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
    handleOptionChange(val) {
      this.multipleOption = val
      this.changeUser()
    },
    async kpiDataEnter(){
        const ids = []
        const score = []
        const nickName =[]
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
              var userData =item.EvaluationKpis[i].Users
              score.push(data)
              for (let u = 0; u < userData.length; u++) {
                nickName.push(userData[u].nickName)
              }
            }
          })
          const checkArr = await getKpiByIds({ ids })
          //const usersArr =  await getUserByNickName({nickName})
          const res = await assignedKpiEvaluation({
            kpis: checkArr.data.list,
            id: this.row.ID,
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
      const res = await setUserEvaluation({
        id: this.row.ID,
        userId: ids
      });
      if (res.code == 0) {
        this.$message({ type: "success", message: "角色设置成功" });
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