<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item>
          <el-button @click="openDialog" type="primary" size="mini">添加指标</el-button>
        </el-form-item>
        <el-form-item>
          <el-popover placement="top" v-model="deleteVisible" width="160">
            <p>确定要清空吗？</p>
              <div style="text-align: right; margin: 0">
                <el-button @click="deleteVisible = false" size="mini" type="text">取消</el-button>
                <el-button @click="removeEvaluationKpiByIds" size="mini" type="primary">确定</el-button>
              </div>
            <el-button icon="el-icon-delete" size="mini" slot="reference" type="danger">批量移除</el-button>
          </el-popover>
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
    
    <el-table-column type="selection" width="55"></el-table-column>

    <el-table-column label="指标名称" width="90">
      <template slot-scope="scope">
        <p v-for="(item,index) in scope.row.kpis"
        :key="index">{{item.name}}<br/></p>
      </template>
    </el-table-column>

    <el-table-column label="指标说明" width="360" type="textarea">
      <template slot-scope="scope">
        <span v-for="(item,index) in scope.row.kpis"
        :key="index">{{item.description}}<br/></span>
      </template>
    </el-table-column>

    <el-table-column label="指标算法" width="360" type="textarea">
      <template slot-scope="scope">
        <span v-for="(item,index) in scope.row.kpis"
        :key="index">{{item.category}}<br/></span>
      </template>
    </el-table-column>
    
    <el-table-column label="指标分数" prop="kpiScore" width="80"></el-table-column>
    
    <el-table-column label="评分人" width="230">
      <template slot-scope="scope">
          <el-cascader
            @change="(val)=>{handleOptionChange(val,scope.row)}"
            :v-model="scope.row.Users.ID"
            :options="userOptions"
            :rules="rules"
            clearable
            :props="{ checkStrictly: true,label:'nickName',value:'id',multiple: true,}"
            filterable
          ></el-cascader>
          <span><br/>已选：<br/></span>
          <span v-for="(item,index) in scope.row.Users"
        :key="index">{{item.nickName}}、
        </span>
      </template>
    </el-table-column>
      <el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-button @click="removeKpi(scope.row)" type="danger" size="mini" slot="reference" label="移除指标">移除指标</el-button>
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
    <el-table-column label="指标名称" prop="name" width="120"></el-table-column> 
    
    <el-table-column label="指标说明" prop="description" width="360" type="textarea"></el-table-column> 
    
    <!-- <el-table-column label="指标状态" prop="Status" width="120"></el-table-column>  -->
    
    <el-table-column label="指标算法" prop="category" width="360" type="textarea"></el-table-column> 

     <el-table-column label="指标分数">
      <template slot-scope="scope">
          <el-input v-model="scope.row.evaluationKpis.kpiScore" clearable placeholder="请输入"></el-input>
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
    setUserEvaluation,
    createEvaluationKpi,
    removeEvaluationKpi,
    removeEvaluationKpiByIds,
} from "@/api/pas/evaluationKpi";
import {
    getUserList,
    getUserByIds
} from "@/api/user";
import {
    findEvaluationKpiUser,
} from "@/api/pas/evaluationKpiUser";
import {
    updateEvaluationByInfo
} from "@/api/pas/evaluation";
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
      formData:{
        totalScore:""
      },
      kpiList:{
            name:"",
            ID:"",
            category:"",
            kpiScore:"",
            evaluationKpis:{
              kpiScore:"",
              evaluation_id:"",
              kpi_id:"",
        },
      },
      evaluationData:{
            kcore: "",
      },
      rules: {
        id: [
          { required: true, message: "请选择用户角色", trigger: "blur" }
        ]
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
    async openDialog() {
        const life = await getKpiList({ID:Number(this.row.ID)});
        this.kpiList = life.data.list;
        this.dialogFormVisible = true;
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
     // const arr = this.$steamrollArray(val)
      //console.log(arr,row.ID)
      this.multipleOption = val
      const evaluationKpiId = row.ID
      //this.$message(ID+"123")
      this.changeUser(evaluationKpiId)
    },
    async kpiDataEnter(row){
        const res = await createEvaluationKpi({
        KpiScore: Number(row.evaluationKpis.kpiScore),
        EvaluationId: this.row.ID,
        KpiId: row.ID,
          })
          if(res.code == 0){
              this.$message({ type: 'success', message: "添加成功" })
          }
          this.setTotalScore()
      },
    async refreshEvalutationKpi(){
      const ref = await getKpiEvaluation({ID:Number(this.row.ID)})
        if (ref.code == 0) {
        this.KpiData = ref.data.list;
        }
    },
    async setTotalScore(){
      const ref = await getKpiEvaluation({ID:Number(this.row.ID)})
        this.KpiData = ref.data.list;
        var totalScore = 0
        for (let num = 0; num < this.KpiData.length; num++) {
          totalScore = totalScore + this.KpiData[num].kpiScore
        }
      updateEvaluationByInfo({...this.row,Score:Number(totalScore),ID:Number(this.row.ID)});
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
    async changeUser(evaluationKpiId) {
      const ids = [];
      this.multipleOption &&
          this.multipleOption.map(item => {
            ids.push(Number(item))
          })
      const checkArr = await getUserByIds({ ids })
      const res = await setUserEvaluation({
        id: evaluationKpiId,
        users: checkArr.data.list
      });
      if (res.code == 0) {
        findEvaluationKpiUser({evaluationkpiUserId:evaluationKpiId})
        this.refreshEvalutationKpi()
        this.$message({ type: "success", message: "角色设置成功" });
      }
    },
    
    async removeKpi(row) {
      if (row.Users!=null) {
        await setUserEvaluation({
           id: Number(row.ID),
           users: []
        });
      }
      const res = await removeEvaluationKpi({
        ID: Number(row.ID)
        });
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "移除成功"
        });
        this.setTotalScore()
      }
    },
    async removeEvaluationKpiByIds(){
      const ids = []
        if(this.multipleSelection.length == 0){
          this.$message({
            type: 'warning',
            message: '请选择要删除的数据'
          })
          return
        }
        this.multipleSelection &&
          this.multipleSelection.map(item => {
            ids.push(item.ID)
          })
        const res = await removeEvaluationKpiByIds({ ids })
        if (res.code == 0) {
          this.$message({
            type: 'success',
            message: '删除成功'
          })
        this.setTotalScore()
        this.deleteVisible = false
        }
    },
  },
  async created() {
    //const life = await getKpiScoreByIds({ID:Number(this.row.ID)});
    const user = await getUserList({ page: 1, pageSize: 999 });
    //载入Users
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