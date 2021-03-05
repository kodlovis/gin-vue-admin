<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="考核名称">
          <el-input placeholder="搜索条件" v-model="searchInfo.name"></el-input>
        </el-form-item>    
        <el-form-item label="考核状态">
          <el-select v-model="searchInfo.status" placeholder="请选择" clearable>
            <el-option
              v-for="item in dictList"
              :key="item.value"
              :label="item.label"
              :value="item.value">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="被考核人">
          <el-input placeholder="搜索条件" v-model="searchInfo.nickName"></el-input>
        </el-form-item>  
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
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
    <el-table-column type="selection" width="55"></el-table-column>
    <el-table-column label="考核表名称" prop="name" width="120"></el-table-column>
    <el-table-column label="考核表状态" prop="status" width="120">
      <template slot-scope="scope">
        <!-- <span>{{scope.row.status==0?"评分人确认":""}}</span> -->
        <span>{{filterDict(scope.row.status)}}</span>
      </template>
      </el-table-column>

    <el-table-column label="方案名称" width="120" prop="evaluation.name"></el-table-column>

    <el-table-column label="方案总分" width="120" prop="score"></el-table-column>
    <el-table-column label="方案得分" width="120" prop="result"></el-table-column>

    <el-table-column label="被考核人" width="120" prop="user.nickName"></el-table-column>
    
    <el-table-column label="开始时间" prop="startDate" width="120">
      <template slot-scope="scope">{{scope.row.startDate|formatDate}}</template>
    </el-table-column>

    <el-table-column label="结束时间" prop="endingDate" width="120">
      <template slot-scope="scope">{{scope.row.endingDate|formatDate}}</template>
    </el-table-column>
    
      <el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-button class="table-button" @click="updatePerformancItemReview(scope.row)" size="small" type="primary" >查看指标</el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-pagination
      background
      :current-page="page"
      :page-size="pageSize"
      :page-sizes="[10, 30, 50, 100]"
      :style="{float:'right',padding:'20px'}"
      :total="total"
      @current-change="handleCurrentChange"
      @size-change="handleSizeChange"
      layout="total, sizes, prev, pager, next, jumper"
    ></el-pagination>
    <!-- 编辑指标的弹窗 -->
    <div>
      <el-dialog :before-close="closeprItemDialog" :visible.sync="prItemDialog" title="考核指标" :append-to-body="true" :fullscreen ="true" width="90%">
        <el-table
        :data="performanceReviewItemData"
        @selection-change="handleSelectionChange"
        border
        ref="multipleTable"
        stripe
        style="width: 100%"
        tooltip-effect="dark"
      >
      <el-table-column label="指标名称" prop="kpi.name" width="150"></el-table-column>

      <el-table-column label="指标说明" prop="kpi.description" width="360" type="textarea"></el-table-column>
      
      <el-table-column label="指标算法" prop="kpi.category" width="360" type="textarea"> </el-table-column> 
      <el-table-column label="指标状态" prop="status" width="160">
      <template slot-scope="scope">
        <!-- <span>{{scope.row.status==0?"评分人确认":""}}</span> -->
        <span>{{filterKpiDict(scope.row.status)}}</span>
      </template></el-table-column>
      <el-table-column label="指标分值" prop="score" width="100"></el-table-column> 
      <el-table-column label="指标得分" prop="result" width="100"></el-table-column>
      <el-table-column label="评分人打分详情">
        <template slot-scope="scope">
          <span v-for="(item,index) in scope.row.PRIUs"
          :key="index">评分人：{{item.user.nickName}},评分：{{item.result}}分<br/></span>
        </template>
      </el-table-column>
      </el-table>
      <el-pagination
      background
        :current-page="priPage"
        :page-size="priPageSize"
        :page-sizes="[5,10, 30, 50, 100]"
        :style="{float:'right',padding:'20px'}"
        :total="priTotal"
        @current-change="kpiCurrentChange"
        @size-change="kpiSizeChange"
        layout="total, sizes, prev, pager, next, jumper"
      ></el-pagination>
      </el-dialog>
    </div>
    <el-dialog :before-close="closeRatioDialog" :visible.sync="ratioDialog" title="编辑人员权重" :append-to-body="true" style="width: 50%,marigin:right"
     >
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item>
          <el-button @click="openUserDialog" type="primary" size="mini">添加评分人</el-button>
        </el-form-item>
      </el-form>
    <el-table
      :data="priuData"
      @selection-change="handleSelectionChange"
      border
      ref="multipleTable"
      stripe
      style="width: 80%"
      tooltip-effect="dark"
    >
    <el-table-column label="评分人" width="140">
      <template slot-scope="scope">
          <el-cascader
            @change="(val)=>{handleOptionChange(val,scope.row)}"
            v-model="scope.row.userId"
            :options="userOptions"
            :rules="rules"
            clearable
            :props="{ checkStrictly: true,label:'nickName',value:'id'}"
            filterable
          ></el-cascader>
      </template>
    </el-table-column> 
      <el-table-column label="操作状态" prop="kpi.status" width="160">
        <template slot-scope="scope">
        <el-select v-model="scope.row.status" placeholder="请选择" clearable>
          <el-option
            v-for="item in kpiDictList"
            :key="item.value"
            :label="item.label"
            :value="item.value">
          </el-option>
        </el-select>
        </template>
      </el-table-column>
      <el-table-column label="当前得分" width="160">
        <template slot-scope="scope">
            <el-input v-model="scope.row.result" clearable placeholder="请输入"></el-input>
        </template>
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
              <el-button type="danger" size="mini" slot="reference" label="移除评分人" :disabled="isDisable">移除评分人</el-button>
            </el-popover>
        </template>
      </el-table-column>
    </el-table>
    </el-dialog>
  </div>
</template>

<script>
import {
    deletePerformanceReview,
    findPerformanceReview,
    getPerformanceReviewList,
    getPRByID
} from "@/api/pas/performanceReview";
import {
    getPRIUByPRIID,
} from "@/api/pas/performanceReviewItemUser";
import {
    getEvaluationList,
} from "@/api/pas/evaluation";
import {
    getUserList,
    getUserListByAuthorityId
} from "@/api/user";
import {
    deletePerformanceReviewItem,
    getPerformanceReviewItemListById,
} from "@/api/pas/performanceReviewItem";
import {
    getKpiList,
} from "@/api/pas/kpi";  //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/date";
import { getDict } from "@/utils/dictionary";
import infoList from "@/mixins/infoList";
export default {
  name: "PerformanceReview",
  mixins: [infoList],
  data() {
    return {
      listApi: getPerformanceReviewList,
      dialogFormVisible: false,
      kpiDialog:false,
      prItemDialog: false,
      userDialog: false,
      ratioDialog: false,
      visible: false,
      isDisable:false,
      dialogFlag: false,
      loading:false,
      saveID:0,
      page: 1,
      PRIPage: 1,
      priPage: 1,
      total: 10,
      pageSize: 10,
      PRIPageSize: 10,
      priPageSize: 10,
      type: "",
      searchInfo: {},
      priTotal:0,
      PRITotal:0,
      kpiDictList:[],
      userOptions:[],
      evaluationOptions:[],
      dictList:[],
      deleteVisible: false,
      confirmVisible: false,
      Selection: "",
      totalScore:0,
      multipleSelection: [],formData: {
            id:"",
            name:"",
            status:"",
            startDate:new Date(),
            endingDate:new Date(),
            evaluationId:"",
            employeeId:"",
            score:0,
      },
      kpiList:{
            name:"",
            ID:"",
            category:"",
            kpiScore:"",
            userId:"",
            evaluationKpis:{
              kpiScore:"",
        },
      },
      evaluationData:{
            ID:"",
            name:"",
            category:"",
            status:"",
            description:"",
            score: "",
            Kpis:0,
      },
      priuData:{},
      lastPRData:{},
      saveData:{
        priid:"",
      },
      performanceReviewItemData:{
            PRId:"",
            kpiId:"",
            userId:"",
            score:"",
            kpi:{
              name:"",
              description:"",
              category:"",
            },
            user:{
              ID:"",
            }
      },
    };
  },
  filters: {
    formatDate: function(time) {
      if (time != null && time != "") {
        var date = new Date(time);
        return formatTimeToStr(date, "yyyy-MM-dd");
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
      async openKpiDialog(){
        this.kpiDialog = true;
        const ref = await getKpiList({
          page: this.page, 
          pageSize: this.pageSize})
        if (ref.code == 0) {
        this.kpiList = ref.data.list;
        this.kpiTotal=ref.data.total
        }
      },
      filterDict(status){
        const re = this.dictList.filter(item=>{
          return item.value == status
        })[0]
        if(re){
          return re.label
          }
        else{
          return""
          }
      },
      filterKpiDict(status){
        const re = this.kpiDictList.filter(item=>{
          return item.value == status
        })[0]
        if(re){
          return re.label
          }
        else{
          return""
          }
      },
    handleOptionChange(val){
      this.Selection = val
    },
    setEvaluationOptions(evaluationData) {
      this.evaluationOptions = [];
      this.ids = [];
      this.setEvaluationrOptionsData(evaluationData, this.evaluationOptions ,this.ids);
    },
    setEvaluationrOptionsData(evaluationData, optionsData ,ids) {
      evaluationData &&
        evaluationData.map(item => {
            const option = {
              id: item.ID,
              name: item.name
            };
            optionsData.push(option);
            const idOption = {
              id: item.ID,
            };
            ids.push(idOption)
        });
    },
    setUserOptions(userData) {
      this.userOptions = [];
      this.ids = [];
      this.setUserOptionsData(userData, this.userOptions ,this.ids);
    },
    setUserOptionsData(UserData, optionsData ,ids) {
      UserData &&
        UserData.map(item => {
            const option = {
              id: item.ID,
              nickName: item.nickName
            };
            optionsData.push(option);
            const idOption = {
              id: item.ID,
            };
            ids.push(idOption)
        });
    },
      onSubmit() {
        this.page = 1
        this.pageSize = 10           
        this.getTableData()
      },
      async searchPRIInfo() {
        this.PRIPage = 1
        this.PRIPageSize = 10
        const ref = await getPerformanceReviewItemListById({...this.searchInfo,
          page: this.PRIPage, 
          pageSize: this.PRIPageSize})
        this.PRITotal=ref.data.total
        this.PRIList=ref.data.list
      },
      handleSelectionChange(val) {
        this.multipleSelection = val
      },
    async updatePerformanceReview(row) {
      const res = await findPerformanceReview({ ID: row.ID });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.rePerformanceReview;
        this.isDisable=true
        this.dialogFormVisible = true;
      }
    },
    //打开编辑指标的弹窗并传参
    async updatePerformancItemReview(row) {
      const res = await getPerformanceReviewItemListById({ PRId: row.ID,
        page: this.priPage, 
        pageSize: this.priPageSize})
      this.priTotal=res.data.total
      this.saveID = row.ID
      const pr = await getPRByID({ ID: row.ID,
          page: this.page, 
          pageSize: this.pageSize });
      this.formData = pr.data.list[0];
      if (res.code == 0) {
        this.performanceReviewItemData = res.data.list;
        this.prItemDialog = true;
      }
    },
    async openUserDialog(){
        const res = await getUserListByAuthorityId({authorityIds:["10000","8888"]})
        this.userData = res.data.list
        this.userDialog = true
    },
    async openRatioDialog(row) {
        const res = await getPRIUByPRIID({priid:row.ID})
        this.priuData = res.data.list
        this.saveData.priid=row.ID
        this.ratioDialog = true
      },
    closeRatioDialog() {
      this.ratioDialog = false;
    },
    closeUserDialog() {
      this.userDialog = false;
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.isDisable=false
      this.formData = {
          StartDate:new Date(),
          EndingDate:new Date(),
      };
    },
    closeprItemDialog() {
      this.prItemDialog = false;
    },
    closeKpiDialog(){
      this.getTableData()
      this.kpiDialog = false;
    },
    async kpiSizeChange(val) {
        this.priPageSize = val
        const res = await getPerformanceReviewItemListById({ PRId: this.saveID,
          page: this.priPage, 
          pageSize: this.priPageSize})
        this.priTotal=res.data.total
        this.performanceReviewItemData = res.data.list
    },
    async kpiCurrentChange(val) {
        this.priPage = val
        const res = await getPerformanceReviewItemListById({ PRId: this.saveID,
          page: this.priPage, 
          pageSize: this.priPageSize})
        this.priTotal=res.data.total
        this.performanceReviewItemData = res.data.list
    },
    async kpiStoreSizeChange(val) {
        this.kpiPageSize = val
        const ref = await getKpiList({
          page: this.kpiPage, 
          pageSize: this.kpiPageSize})
        this.kpiTotal=ref.data.total
        this.kpiList=ref.data.list
    },
    async kpiStoreCurrentChange(val) {
        this.kpiPage = val
        const ref = await getKpiList({
          page: this.kpiPage, 
          pageSize: this.kpiPageSize})
        this.kpiTotal=ref.data.total
        this.kpiList=ref.data.list
    },
    async deletePerformanceReview(row) {
      this.isDisable=true;
      const res = await deletePerformanceReview({ ID: row.ID });
      deletePerformanceReviewItem({ PRId: row.ID })
      this.visible = false;
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "删除成功"
        });
        this.isDisable=false
        this.getTableData();
      }
    },
    async openDialog() {
      this.type = "create";
      this.dialogFormVisible = true;
    }
  },
  async created() {
    //获取考核状态字典
    const pr = await getDict("PR");
    pr.map(item=>item.value)
    this.dictList = pr
    await this.getTableData();
    //载入Users
    const user = await getUserList({ page: 1, pageSize: 999 });
    this.setUserOptions(user.data.list);
    //获取指标状态字典
    const kpi = await getDict("kpi");
    kpi.map(item=>item.value)
    this.kpiDictList = kpi
    //载入Evaluations
    const evaluations = await getEvaluationList({ page: 1, pageSize: 999 });
    this.setEvaluationOptions(evaluations.data.list);
}
};
</script>

<style>
</style>