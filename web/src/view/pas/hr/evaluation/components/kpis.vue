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
                <el-button @click="removeEvaluationKpiByIds" size="mini" type="primary" :disabled="isDisable">确定</el-button>
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
        <span v-for="(item,index) in scope.row.evaluationKpiUsers"
        :key="index">{{item.user.nickName}}<br/></span>
      </template>
    </el-table-column>
      <el-table-column label="按钮组" width="230">
        <template slot-scope="scope">
          <el-popover placement="top" width="160" v-model="scope.row.visible">
              <p>确定要删除吗？</p>
              <div style="text-align: right; margin: 0">
                <el-button size="mini" type="text" @click="scope.row.visible = false">取消</el-button>
                <el-button type="primary" size="mini" @click="removeKpi(scope.row)" :disabled="isDisable">确定</el-button>
              </div>
              <el-button type="danger" icon="el-icon-delete" size="mini" slot="reference" label="移除指标" :disabled="isDisable">移除指标</el-button>
              <el-button @click="openRatioDialog(scope.row)" type="primary" size="mini" slot="reference" label="编辑人员权重">编辑人员权重</el-button>
            </el-popover>
        </template>
      </el-table-column>
    </el-table>

    <el-pagination
      :current-page="page"
      :page-size="pageSize"
      :page-sizes="[5,10,15,20]"
      :style="{float:'right',padding:'20px'}"
      :total="total"
      @current-change="handleCurrentChange"
      @size-change="handleSizeChange"
      layout="total, sizes, prev, pager, next, jumper"
    ></el-pagination>

    <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="添加指标" :append-to-body="true" style="width: 90%,marigin:right"
     >
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
    <el-table-column label="设置指标分数">
      <template slot-scope="scope">
          <el-input v-model="scope.row.evaluationKpis.kpiScore" clearable placeholder="请输入"></el-input>
      </template>
    </el-table-column>
    <el-table-column label="设置评分人" width="230">
      <template slot-scope="scope">
          <el-cascader
            @change="(val)=>{handleOptionChange(val,scope.row)}"
            v-model="scope.row.userId"
            :options="userOptions"
            :rules="rules"
            clearable
            :props="{ checkStrictly: true,label:'nickName',value:'id',multiple: true }"
            filterable
          ></el-cascader>
      </template>
    </el-table-column>
      <el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-button @click="kpiDataEnter(scope.row)" type="primary" size="mini" slot="reference" label="添加" :disabled="isDisable">添加</el-button>
        </template>
      </el-table-column>
    </el-table>
    </el-dialog>

    <el-dialog :before-close="closeRatioDialog" :visible.sync="ratioDialog" title="编辑人员权重" :append-to-body="true" style="width: 50%,marigin:right"
     >
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item>
          <el-button @click="openUserDialog" type="primary" size="mini">添加评分人</el-button>
        </el-form-item>
      </el-form>
    <el-table
      :data="ekuData"
      @selection-change="handleSelectionChange"
      border
      ref="multipleTable"
      stripe
      style="width: 50%"
      tooltip-effect="dark"
    >
    <el-table-column label="评分人" prop="user.nickName" width="120">
    </el-table-column> 

    <el-table-column label="设置权重">
        <template slot-scope="scope">
          <el-input v-model="scope.row.score" clearable placeholder="请输入"></el-input>
        </template>
    </el-table-column>
      <el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-button @click="changeRatioData(scope.row)" type="primary" size="mini" slot="reference" label="修改" :disabled="isDisable">修改</el-button>
          <el-popover placement="top" width="160" v-model="scope.row.visible">
              <p>确定要删除吗？</p>
              <div style="text-align: right; margin: 0">
                <el-button size="mini" type="text" @click="scope.row.visible = false">取消</el-button>
                <el-button type="primary" size="mini" @click="removeUser(scope.row)" :disabled="isDisable">确定</el-button>
              </div>
              <el-button type="danger" icon="el-icon-delete" size="mini" slot="reference" label="移除评分人" :disabled="isDisable">移除评分人</el-button>
            </el-popover>
        </template>
      </el-table-column>
    </el-table>
    </el-dialog>
    <el-dialog :before-close="closeUserDialog" :visible.sync="userDialog" title="增加评分人" :append-to-body="true" style="width: 50%,marigin:right"
     >
    <el-table
      :data="userData"
      @selection-change="handleSelectionChange"
      border
      ref="multipleTable"
      stripe
      style="width: 50%"
      tooltip-effect="dark"
    >
    <el-table-column label="评分人" prop="nickName" width="120">
    </el-table-column> 
      <el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-button @click="userDataEnter(scope.row)" type="primary" size="mini" slot="reference" label="增加" :disabled="isDisable">增加</el-button>
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
    createEvaluationKpi,
    removeEvaluationKpi,
    removeEvaluationKpiByIds,
    getLastEvaluationKpi
} from "@/api/pas/evaluationKpi";
import {
    getUserList,
    getUserListByAuthorityId
} from "@/api/user";
import {
    createEKU,
    getEKUByEKID,
    updateEKU,
    removeEKU
} from "@/api/pas/evaluationKpiUser";
import {
    updateEvaluationByInfo,
    findEvaluation,
    updateEvaluation
} from "@/api/pas/evaluation";
import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";
export default {
  name: "kpis",
  mixins: [infoList],
  listApi: getKpiEvaluation,
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
      dialogFormVisible: false,
      ratioDialog: false,
      userDialog: false,
      visible: false,
      loading: false,
      ids:[],
      page: 1,
      total: 10,
      pageSize: 5,
      isDisable:false,
      type: "",
      deleteVisible: false,
      multipleOption: [],
      multipleSelection: [],
      formData:{
        totalScore:"",
      },
      KpiData: {
        evaluationKpiUsers:[{
          user:{
            nickName:"",
          }
        }]
      },
      saveData:{
        ekid:0,
      },
      kpiList:{
            name:"",
            ID:"",
            category:"",
            kpiScore:"",
            userId:"",
            evaluationKpis:{
              evaluationKpiUsers:[{
                user:{
                  nickName:"",
                }
              }],
              kpiScore:"",
              evaluation_id:"",
              kpi_id:"",
        },
      },
      evaluationData:{
            score: "",
      },
      ekData:{

      },
      ekuData:{
        user:{
        nickName:"",
        },
        score:"",
      },
      rules: {
        id: [
          { required: true, message: "请选择用户角色", trigger: "blur" }
        ],
        nickName:{ required: true, message: "请选择用户角色", trigger: "blur" }
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
    async userDataEnter(row){
      var items=[{
        ekid:this.saveData.ekid,
        userid:row.ID,
        score:Number((Number(1)/Number(this.ekuData.length+1)).toFixed(5))
      }]
      const ref = await createEKU({
        items:items
      })
      if(ref.code == 0){
        const res = await getEKUByEKID({ekid:this.saveData.ekid,authorityId:"10000"})
        this.ekuData = res.data.list
        this.$message({
          type:"success",
          message:"添加成功"
        })
      }
    },
    async changeRatioData(row){
      this.isDisable=true;
      if(row.score<=0||row.score>1){
        this.$message({
          type:"success",
          message:"无效输入"
        })
          this.isDisable=false;
          return
        }
      const ref = await updateEKU({...row,
        score:Number(row.score),
      })
        if (ref.code == 0) {
          this.$message({
            type:"success",
            message:"添加成功"
          })
          this.isDisable=false;
          const res = await getEKUByEKID({ekid:row.ekid,authorityId:"10000"})
          this.ekuData = res.data.list
        }
    },
    async openRatioDialog(row) {
        const res = await getEKUByEKID({ekid:row.ID,authorityId:"10000"})
        this.ekuData = res.data.list
        this.saveData.ekid=row.ID
        this.ratioDialog = true
      },
    async openUserDialog(){
        const res = await getUserListByAuthorityId({authorityId:"10000"})
        this.userData = res.data.list
        this.userDialog = true
    },
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
      //const evaluationKpiId = row.ID
      const userId = row.ID
      // this.changeUser(evaluationKpiId)
      return userId
    },
    async kpiDataEnter(row){
        this.isDisable=true;
        const res = await createEvaluationKpi({
          EvaluationId: this.row.ID,
          KpiId: row.ID,
          kpiScore:Number(row.evaluationKpis.kpiScore),
        })
          if(res.code == 0){
          var ek = await getLastEvaluationKpi()
          this.KpiData = ek.data.reEK
          var items= []
          for (let i = 0; i < row.userId.length; i++) {
            items.push({
              userId:Number(row.userId[i]),
              ekid:Number(this.KpiData.ID),
              score:Number((Number(1)/Number(row.userId.length)).toFixed(5)),   
            })
          }
            if(ek.code == 0){
              createEKU({
                items:items
              })
            this.isDisable=false
            //计算更新当条方案分数
            var total = Number(this.KpiData.evaluation.score)+Number(row.evaluationKpis.kpiScore)
            this.KpiData = this.KpiData.evaluation
            const ud = await updateEvaluation({...this.KpiData,score:total})
            if(ud.code == 0){
              this.refreshEvalutationKpi(),
              this.$message({ type: 'success', message: "添加成功" })
              }
            }
          }
      },
    async refreshEvalutationKpi(page = this.page, pageSize = this.pageSize){
      const ref = await getKpiEvaluation({
            ID:Number(this.row.ID),
            page: page, 
            pageSize: pageSize
          })
            this.KpiData = ref.data.list;
            this.total = ref.data.total
            this.page = ref.data.page
            this.pageSize = ref.data.pageSize
    },
    handleSizeChange(val) {
        this.pageSize = val
        this.refreshEvalutationKpi()
    },
    handleCurrentChange(val) {
        this.page = val
        this.refreshEvalutationKpi()
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
          name:"",
          description:"",
          status:"",
          category:"",
          tags: 1,
      };
    },
    closeRatioDialog() {
      this.ratioDialog = false;
    },
    closeUserDialog() {
      this.userDialog = false;
      this.refreshEvalutationKpi()
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
    // async changeUser(evaluationKpiId) {
    //   const ids = [];
    //   this.multipleOption &&
    //       this.multipleOption.map(item => {
    //         ids.push(Number(item))
    //       })
    //   const checkArr = await getUserByIds({ ids })
    //   const res = await setUserEvaluation({
    //     id: evaluationKpiId,
    //     user: checkArr.data.list
    //   });
    //   if (res.code == 0) {
    //     findEvaluationKpiUser({evaluationkpiUserId:evaluationKpiId})
    //     this.refreshEvalutationKpi()
    //     this.$message({ type: "success", message: "角色设置成功" });
    //   }
    // },
    
    async removeKpi(row) {
      // if (row.Users!=null) {
      //   await setUserEvaluation({
      //      id: Number(row.ID),
      //      users: []
      //   });
      //删除方案中的指标
      this.isDisable=true;
      row.visible = false;
      removeEvaluationKpi({
        ID: Number(row.ID)
        });
        const ref = await findEvaluation({ID:Number(this.row.ID)})
        if (ref.code == 0) {
          this.evaluationData = ref.data.reEvaluation;
          var totalScore = 0
          totalScore = this.evaluationData.score - row.kpiScore
          updateEvaluationByInfo({...this.row,score:Number(totalScore),ID:Number(this.row.ID)});
        }
        this.refreshEvalutationKpi()
        this.$message({
          type: "success",
          message: "移除成功"
        });
      this.isDisable=false
    },
    async removeUser(row){
      this.isDisable=true;
      row.visible = false;
      const ref = await removeEKU({ID:row.id,ekid:row.ekid})
      if(ref.code ==0){
        this.$message({
          type: "success",
          message: "移除成功"
        });
        this.isDisable=false;
        const res = await getEKUByEKID({ekid:row.ekid,authorityId:"10000"})
        this.ekuData = res.data.list
      }
    },
    async removeEvaluationKpiByIds(){
      this.isDisable=true;
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
        this.deleteVisible = false
        this.isDisable=false
        if (res.code == 0) {
          //拿本方案的总分一次次减去选中的分数
          const ref = await findEvaluation({ID:Number(this.row.ID)})
          this.evaluationData = ref.data.reEvaluation;
          var sum = Number(this.evaluationData.score)
          for (let i = 0; i < this.multipleSelection.length; i++) {
            sum=sum-this.multipleSelection[i].kpiScore
          }
          //更新方案总分
          updateEvaluationByInfo({...this.row,score:Number(sum),ID:Number(this.row.ID)});
          this.$message({
            type: 'success',
            message: '删除成功'
          })
          this.refreshEvalutationKpi()
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