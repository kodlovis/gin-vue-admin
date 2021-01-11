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
        <!-- <el-form-item>
          <el-popover placement="top" v-model="deleteVisible" width="160">
            <p>确定要删除吗？</p>
              <div style="text-align: right; margin: 0">
                <el-button @click="deleteVisible = false" size="mini" type="text">取消</el-button>
                <el-button @click="onDelete" size="mini" type="primary">确定</el-button>
              </div>
            <el-button icon="el-icon-delete" size="mini" slot="reference" type="danger">批量删除</el-button>
          </el-popover>
        </el-form-item> -->
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
    
    <!-- <el-table-column label="方案类型" prop="Category" width="120"></el-table-column>  -->
    
    <el-table-column label="考核表状态" prop="status" width="120"></el-table-column>

    <el-table-column label="方案名称" width="120" prop="evaluation.name"></el-table-column>

    <el-table-column label="方案总分" width="120" prop="evaluation.score"></el-table-column>

    <el-table-column label="被考核人" width="120" prop="user.nickName"></el-table-column>
    
    <el-table-column label="开始时间" prop="startDate" width="120"></el-table-column>

    <el-table-column label="结束时间" prop="endingDate" width="120"></el-table-column>
    
      <el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-button class="table-button" @click="updatePerformanceReview(scope.row)" size="small" type="primary" icon="el-icon-edit">变更</el-button>
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

    <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="弹窗操作">
      <el-form :model="formData" label-position="right" label-width="120px">
       
         <el-form-item label="考核表名称:">
            <el-input v-model="formData.name" clearable placeholder="请输入" :style="{width: '60%'}"></el-input>
      </el-form-item>
       
         <el-form-item label="考核表状态:">
            <el-input v-model="formData.status" clearable placeholder="请输入" :style="{width: '60%'}"></el-input>
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
    <!-- <el-dialog :before-close="closeEvaluationDialog" :visible.sync="dialogEvaluationForm" title="弹窗操作">
      <el-table
      :data="evaluationData"
      @selection-change="handleSelectionChange"
      border
      ref="multipleTable"
      stripe
      style="width: 100%"
      tooltip-effect="dark"
    >
    <el-table-column label="方案名称" prop="name" width="120"></el-table-column> 
    
    <el-table-column label="方案类型" prop="Category" width="120"></el-table-column>
    
    <el-table-column label="方案状态" prop="status" width="120"></el-table-column> 
    
    <el-table-column label="方案描述" prop="description" width="120"></el-table-column> 
    
    <el-table-column label="方案总分" prop="score" width="120"></el-table-column> 
    
      <el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-button class="table-button" @click="chooseEvaluation(scope.row)" size="small" type="primary" icon="el-icon-edit">添加</el-button>
        </template>
      </el-table-column>
    </el-table>
    </el-dialog> -->

  </div>
</template>

<script>
import {
    createPerformanceReview,
    deletePerformanceReview,
    deletePerformanceReviewByIds,
    findPerformanceReview,
    getPerformanceReviewList,
    updatePerformanceReviewByInfo
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
    createPerformanceReviewItem
} from "@/api/pas/performanceReviewItem";
import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";
export default {
  name: "PerformanceReview",
  mixins: [infoList],
  data() {
    return {
      listApi: getPerformanceReviewList,
      dialogFormVisible: false,
      visible: false,
      type: "",
      deleteVisible: false,
      Selection: "",
      totalScore:0,
      multipleSelection: [],formData: {
            name:"",
            status:"",
            startDate:new Date(),
            endingDate:new Date(),
            evaluationId:"",
            employeeId:"",
            score:0,
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
      evaluationKpiData:{
            ID:"",
            kpiId:"",
            kpiScore:"",
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
              name: item.name
            };
            optionsData.push(option);
            const idOption = {
              id: item.ID,
            };
            ids.push(idOption)
          //}
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
  //条件搜索前端看此方法
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
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
          StartDate:new Date(),
          EndingDate:new Date(),
          
      };
    },
    async deletePerformanceReview(row) {
      this.visible = false;
      const res = await deletePerformanceReview({ ID: row.ID });
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "删除成功"
        });
        this.getTableData();
      }
    },
    async createPerformanceReviewItem(){
        const ref = await getEvaluationKpiById({ID:Number(this.formData.evaluationId)});
        this.evaluationKpiData = ref.data.list
        const item = []
        for (let i = 0; i < this.evaluationKpiData.length; i++) {
          item.push({
            score:Number(this.evaluationKpiData[i].kpiScore),
            kpiId:Number(this.evaluationKpiData[i].kpiId),
            userId:Number(this.evaluationKpiData[i].userId),
            PRId:Number(this.evaluationKpiData[i].ID)
            })
        } 
        createPerformanceReviewItem({item})
        return this.totalScore
    },
    async enterDialog() {
      let res;
      switch (this.type) {
        case "create":
          createPerformanceReviewItem()
          var ref = await findEvaluation({ID:Number(this.formData.evaluationId)})
          var evaluationData = ref.data.reEvaluation;
          res = await createPerformanceReview({...this.formData,evaluationId:Number(this.formData.evaluationId),employeeId:Number(this.formData.employeeId),score:Number(evaluationData.score)});
          break;
        case "update":
          res = await updatePerformanceReviewByInfo({...this.formData,evaluationId:Number(this.formData.evaluationId),employeeId:Number(this.formData.employeeId)});
          break;
        default:
          createPerformanceReviewItem()
          res = await createPerformanceReview(this.formData);
          break;
      }
      if (res.code == 0) {
        //  this.$message({
        //   type:"success",
        //   message:"创建/更改成功"
        // })
        this.closeDialog();
        this.getTableData();
      }
    },
    openDialog() {
      this.type = "create";
      this.dialogFormVisible = true;
    }
  },
  async created() {
    await this.getTableData();
    const evaluations = await getEvaluationList({ page: 1, pageSize: 999 });
    //载入Evaluations
    this.setEvaluationOptions(evaluations.data.list);
    const user = await getUserList({ page: 1, pageSize: 999 });
    //载入Users
    this.setUserOptions(user.data.list);
}
};
</script>

<style>
</style>