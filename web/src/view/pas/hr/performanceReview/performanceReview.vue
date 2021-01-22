<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">              
        <!-- <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item> -->
        <el-form-item>
          <el-button @click="openDialog" type="primary">创建考核表</el-button>
        </el-form-item>
        <el-form-item>
          <el-popover placement="top" v-model="deleteVisible" width="160">
            <p>确定要删除吗？</p>
              <div style="text-align: right; margin: 0">
                <el-button @click="deleteVisible = false" size="mini" type="text">取消</el-button>
                <el-button @click="onDelete" size="mini" type="primary">确定</el-button>
              </div>
            <el-button icon="el-icon-delete" size="mini" slot="reference" type="danger">批量删除</el-button>
          </el-popover>
        </el-form-item>
        <el-form-item>
          <el-popover placement="top" v-model="confirmVisible" width="160">
            <p>要开始确认考核吗？</p>
              <div style="text-align: right; margin: 0">
                <el-button @click="confirmVisible = false" size="mini" type="text">取消</el-button>
                <el-button @click="onConfirm" size="mini" type="primary">确定</el-button>
              </div>
            <el-button size="mini" slot="reference" type="primary">开始确认考核</el-button>
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

    <el-table-column label="被考核人" width="120" prop="user.nickName"></el-table-column>
    
    <el-table-column label="开始时间" prop="startDate" width="120"></el-table-column>

    <el-table-column label="结束时间" prop="endingDate" width="120"></el-table-column>
    
      <el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-button class="table-button" @click="updatePerformanceReview(scope.row)" size="small" type="primary" icon="el-icon-edit">编辑考核表</el-button>
          <el-button class="table-button" @click="updatePerformancItemReview(scope.row)" size="small" type="primary" icon="el-icon-edit">编辑指标</el-button>
          <el-popover placement="top" width="160" v-model="scope.row.visible">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin: 0">
              <el-button size="mini" type="text" @click="scope.row.visible = false">取消</el-button>
              <el-button type="primary" size="mini" @click="deletePerformanceReview(scope.row)">确定</el-button>
            </div>
            <el-button type="danger" icon="el-icon-delete" size="mini" slot="reference">删除</el-button>
          </el-popover>
        </template>
      </el-table-column>
    </el-table>

    <el-pagination
      :current-page="page"
      :page-size="pageSize"
      :page-sizes="[10, 30, 50, 100]"
      :style="{float:'right',padding:'20px'}"
      :total="total"
      @current-change="handleCurrentChange"
      @size-change="handleSizeChange"
      layout="total, sizes, prev, pager, next, jumper"
    ></el-pagination>
    <!-- 编辑考核表的弹窗 -->
    <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="编辑考核表">
      <el-form :model="formData" label-position="right" label-width="120px">
       
         <el-form-item label="考核表名称:">
            <el-input v-model="formData.name" clearable placeholder="请输入" :style="{width: '60%'}"></el-input>
      </el-form-item>
       
         <el-form-item label="考核表状态:">
          <el-select v-model="formData.status" placeholder="请选择" clearable>
            <el-option
              v-for="item in dictList"
              :key="item.value"
              :label="item.label"
              :value="item.value">
              </el-option>
            </el-select>  
        </el-form-item>
       
         <el-form-item label="方案选择:" v-model="formData.evaluationId">
           <el-cascader
             @change="(val)=>{handleOptionChange(val)}"
             v-model="formData.evaluationId"
             :options="evaluationOptions"
             :rules="rules"
             clearable
             :props="{ checkStrictly: true,label:'name',value:'id',}"
             filterable
          ></el-cascader>
      </el-form-item>
          <el-form-item label="被考核人选择:">
           <el-cascader
             @change="(val)=>{handleOptionChange(val)}"
             v-model="formData.employeeId"
             :options="userOptions"
             :rules="rules"
             clearable
             :props="{ checkStrictly: true,label:'nickName',value:'id',}"
          ></el-cascader>
      </el-form-item>
         <el-form-item label="开始日期:">
              <el-date-picker type="date" placeholder="选择日期" v-model="formData.startDate" clearable></el-date-picker>
       </el-form-item>
       
         <el-form-item label="结束日期:">
              <el-date-picker type="date" placeholder="选择日期" v-model="formData.endingDate" clearable></el-date-picker>
       </el-form-item>
       </el-form>
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button @click="enterDialog" type="primary">确 定</el-button>
      </div>
    </el-dialog>
    <!-- 编辑指标的弹窗 -->
    <el-dialog :before-close="closeprItemDialog" :visible.sync="prItemDialog" title="编辑指标" :append-to-body="true" width="80%">
      <el-table
      :data="performanceReviewItemData"
      @selection-change="handleSelectionChange"
      border
      ref="multipleTable"
      stripe
      style="width: 100%"
      tooltip-effect="dark"
    >
    <el-table-column label="指标名称" prop="kpi.name" width="90"></el-table-column>

    <el-table-column label="指标说明" prop="kpi.description" width="360" type="textarea"></el-table-column>
    
    <el-table-column label="指标算法" prop="kpi.category" width="360" type="textarea"> </el-table-column> 
    <el-table-column label="指标状态" prop="kpi.status" width="160">
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
    <el-table-column label="指标分数">
      <template slot-scope="scope">
          <el-input v-model="scope.row.score" clearable placeholder="请输入"></el-input>
      </template>
    </el-table-column> 
    <el-table-column label="评分人" width="230">
      <template slot-scope="scope">
          <el-cascader
            @change="(val)=>{handleOptionChange(val,scope.row)}"
            v-model="scope.row.user.ID"
            :options="userOptions"
            :rules="rules"
            clearable
            :props="{ checkStrictly: true,label:'nickName',value:'id',}"
            filterable
          ></el-cascader>
      </template>
    </el-table-column>
      <el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-button @click="performanceReviewItemDataEnter(scope.row)" type="primary" size="mini" slot="reference" label="修改">修改</el-button>
        </template>
      </el-table-column>
    </el-table>
    </el-dialog>

  </div>
</template>

<script>
import {
    createPerformanceReview,
    deletePerformanceReview,
    deletePerformanceReviewByIds,
    findPerformanceReview,
    getPerformanceReviewList,
    updatePerformanceReviewByInfo,
    getLastPerformanceReview,
    updatePRStatysByIds
} from "@/api/pas/performanceReview";  //  此处请自行替换地址
import {
    getEvaluationList,
    findEvaluation
} from "@/api/pas/evaluation";
import {
    getUserList,
} from "@/api/user";
import {
    getEvaluationKpiById
} from "@/api/pas/evaluationKpi";
import {
    createPerformanceReviewItem,
    deletePerformanceReviewItem,
    deletePerformanceReviewItemByIds,
    getPerformanceReviewItemListById,
    updatePerformanceReviewItemByInfo,
    updatePRItemStatysByIds
} from "@/api/pas/performanceReviewItem";
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
      prItemDialog: false,
      visible: false,
      type: "",
      kpiDictList:[],
      dictList:[],
      deleteVisible: false,
      confirmVisible: false,
      Selection: "",
      totalScore:0,
      multipleSelection: [],formData: {
            id:"",
            name:"",
            status:"",
            dictList:[],
            startDate:new Date(),
            endingDate:new Date(),
            evaluationId:"",
            employeeId:"",
            score:0,
      },
      performanceReviewData:{
            score:0,
            kpi:{
              name:"",
              description:"",
              category:"",
              status:0,
            },
            user:{
              ID:"",
            }
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
      handleSelectionChange(val) {
        this.multipleSelection = val
      },
      async onDelete() {
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
        const res = await deletePerformanceReviewByIds({ ids })
        deletePerformanceReviewItemByIds({ ids })
        if (res.code == 0) {
          this.$message({
            type: 'success',
            message: '删除成功'
          })
          this.deleteVisible = false
          this.getTableData()
        }
      },
    async updatePerformanceReview(row) {
      const res = await findPerformanceReview({ ID: row.ID });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.rePerformanceReview;
        this.dialogFormVisible = true;
      }
    },
    //打开编辑指标的弹窗并传参
    async updatePerformancItemReview(row) {
      const res = await getPerformanceReviewItemListById({ PRId: row.ID });
      const pr = await findPerformanceReview({ ID: row.ID });
      this.formData = pr.data.rePerformanceReview;
      this.type = "update";
      if (res.code == 0) {
        this.performanceReviewItemData = res.data.list;
        this.prItemDialog = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
          StartDate:new Date(),
          EndingDate:new Date(),
      };
    },
    closeprItemDialogg() {
      this.prItemDialog = false;
    },
    async deletePerformanceReview(row) {
      this.visible = false;
      const res = await deletePerformanceReview({ ID: row.ID });
      deletePerformanceReviewItem({ PRId: row.ID })
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "删除成功"
        });
        this.getTableData();
      }
    },
    async performanceReviewItemDataEnter(row){
        const res = await updatePerformanceReviewItemByInfo({
        id:Number(row.ID),
        score:Number(row.score),
        userId:Number(row.user.ID),
        status:Number(row.status),
        });
        const prlist = await getPerformanceReviewItemListById({ PRId: this.formData.ID })
        this.performanceReviewItemData = prlist.data.list
        var totalScore = 0
        for (let num = 0; num < this.performanceReviewItemData.length; num++) {
          totalScore = totalScore + this.performanceReviewItemData[num].score
        }
        const ref = await updatePerformanceReviewByInfo({...this.formData,
        evaluationId:Number(this.formData.evaluationId),
        score:Number(totalScore),
        employeeId:Number(this.formData.employeeId)});
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "修改成功"
        });
        if (ref.code == 0){
        this.getTableData();
        }
      }
    },
    async enterDialog() {
      let res;
      switch (this.type) {
        case "create":
          var ref = await findEvaluation({ID:Number(this.formData.evaluationId)})
          var evaluationData = ref.data.reEvaluation;
          res = await createPerformanceReview({...this.formData,
            evaluationId:Number(this.formData.evaluationId),
            employeeId:Number(this.formData.employeeId),
            score:Number(evaluationData.score),
            status:Number(this.formData.status),
            });
          //查询出ID，循环将默认数据组合转入一个list
          var evaluationKpi = await getEvaluationKpiById({ID:Number(this.formData.evaluationId)});
          this.evaluationKpiData = evaluationKpi.data.list
          var pr = await getLastPerformanceReview()
          this.formData = pr.data.rePR
          var item = []
          for (let i = 0; i < this.evaluationKpiData.length; i++) {
            item.push({
              score:Number(this.evaluationKpiData[i].kpiScore),
              kpiId:Number(this.evaluationKpiData[i].kpiId),
              userId:Number(this.evaluationKpiData[i].userId),
              PRId:Number(this.formData.ID),
              status:100,
              })
          } 
        createPerformanceReviewItem({item})
          break;
        case "update":
          res = await updatePerformanceReviewByInfo({...this.formData,
          evaluationId:Number(this.formData.evaluationId),
          status:Number(this.formData.status),
          employeeId:Number(this.formData.employeeId)});
          break;
        default:
          res = await createPerformanceReview(this.formData);
          break;
      }
      if (res.code == 0) {
         this.$message({
          type:"success",
          message:"创建/更改成功"
        })
        this.closeDialog();
        this.getTableData();
      }
    },
    async onConfirm(){
      const ids = []
        if(this.multipleSelection.length == 0){
          this.$message({
            type: 'warning',
            message: '请选择要开始确认的考核'
          })
          return
        }
        if (this.multipleSelection.status!=100) {
          this.$message({
            type: 'warning',
            message: '选中了已开始的考核'
          })
          return
        }
        this.multipleSelection &&
          this.multipleSelection.map(item => {
            ids.push(item.ID)
          })
        const res = await updatePRStatysByIds({ids:ids,status:1})
        if (res.code == 0) {
          this.$message({
            type: 'success',
            message: '发布成功'
          })
          updatePRItemStatysByIds({ids:ids,status:1})
          this.deleteVisible = false
          this.getTableData()
        }
    },
    async openDialog() {
      this.type = "create";
      this.dialogFormVisible = true;
    }
  },
  async created() {
    const res = await getDict("PR");
    res.map(item=>item.value)
    this.dictList = res
    await this.getTableData();
    const user = await getUserList({ page: 1, pageSize: 999 });
    //载入Users
    this.setUserOptions(user.data.list);
    const kpi = await getDict("kpi");
    kpi.map(item=>item.value)
    this.kpiDictList = kpi
    const evaluations = await getEvaluationList({ page: 1, pageSize: 999 });
    //载入Evaluations
    this.setEvaluationOptions(evaluations.data.list);
}
};
</script>

<style>
</style>