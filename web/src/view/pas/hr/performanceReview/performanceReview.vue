<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
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
        <el-form-item label="考核名称">
          <el-input placeholder="搜索条件" v-model="searchInfo.name"></el-input>
        </el-form-item>    
        <el-form-item label="被考核人">
          <el-input placeholder="搜索条件" v-model="searchInfo.nickName"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <br/>
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
    <el-table-column label="开始时间" prop="startDate" width="120">
      <template slot-scope="scope">{{scope.row.startDate|formatDate}}</template>
    </el-table-column>
    <el-table-column label="结束时间" prop="endingDate" width="120">
      <template slot-scope="scope">{{scope.row.endingDate|formatDate}}</template>
    </el-table-column>
      <el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-button class="table-button" @click="updatePerformanceReview(scope.row)" size="small" type="primary" icon="el-icon-edit">编辑考核表</el-button>
          <el-button class="table-button" @click="updatePerformancItemReview(scope.row)" size="small" type="primary" icon="el-icon-edit">编辑指标</el-button>
          <el-popover placement="top" width="160" v-model="scope.row.visible">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin: 0">
              <el-button size="mini" type="text" @click="scope.row.visible = false">取消</el-button>
              <el-button type="primary" size="mini" @click="deletePerformanceReview(scope.row)" :disabled="isDisable">确定</el-button>
            </div>
            <el-button type="danger" icon="el-icon-delete" size="mini" slot="reference">删除</el-button>
          </el-popover>
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
    <!-- 编辑考核表的弹窗 -->
    <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" :title="dialogTitle" 
      v-loading="loading"
      element-loading-text="拼命加载中"
      element-loading-spinner="el-icon-loading"
      element-loading-background="rgba(0, 0, 0, 0.8)">
      <el-form :model="formData" label-position="right" label-width="120px" :rules="rules" ref="prForm">
       
         <el-form-item label="考核表名称:" prop="name">
            <el-input v-model="formData.name" clearable placeholder="请输入" :style="{width: '60%'}"></el-input>
      </el-form-item>
       
         <el-form-item label="考核表状态:" prop="status">
          <el-select v-model="formData.status" placeholder="请选择">
            <el-option
              v-for="item in dictList"
              :key="item.value"
              :label="item.label"
              :value="item.value">
              </el-option>
            </el-select>  
        </el-form-item>
       
         <el-form-item label="方案选择:" v-model="formData.evaluationId" prop="evaluationId">
           <el-cascader
             @change="(val)=>{handleOptionChange(val)}"
             v-model="formData.evaluationId"
             :options="evaluationOptions"
             clearable
             :props="{ checkStrictly: true,label:'name',value:'id',}"
             filterable
            :disabled="isDisable"
          ></el-cascader>
      </el-form-item>
          <el-form-item label="被考核人选择:" prop="employeeId">
           <el-cascader
             @change="(val)=>{handleOptionChange(val)}"
             v-model="formData.employeeId"
             :options="userOptions"
             clearable
             :props="{ checkStrictly: true,label:'nickName',value:'id',}"
             filterable
          ></el-cascader>
      </el-form-item>
         <el-form-item label="开始日期:" prop="startDate">
              <el-date-picker type="date" placeholder="选择日期" v-model="formData.startDate" clearable></el-date-picker>
       </el-form-item>
       
         <el-form-item label="结束日期:" prop="endingDate">
              <el-date-picker type="date" placeholder="选择日期" v-model="formData.endingDate" clearable></el-date-picker>
       </el-form-item>
       </el-form>
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button @click="enterDialog" type="primary">确 定</el-button>
      </div>
    </el-dialog>
    <!-- 编辑指标的弹窗 -->
    <div>
      <el-dialog :before-close="closeprItemDialog" :visible.sync="prItemDialog" title="编辑指标" :append-to-body="true" :fullscreen ="true" width="90%">
        <div class="search-term">
          <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
            <el-form-item>
              <el-button @click="openKpiDialog" type="primary" size="mini">添加指标</el-button>
            </el-form-item>
          </el-form>
        </div>
        <el-table
        :data="performanceReviewItemData"
        @selection-change="handleSelectionChange"
        border
        ref="multipleTable"
        stripe
        style="width: 100%"
        tooltip-effect="dark"
      >
      <el-table-column label="指标类型" width="120">
        <template slot-scope="scope">
            <span v-for="(item,index) in scope.row.kpi.Tags"
        :key="index">{{item.name}}</span>
        </template>
      </el-table-column>
      <el-table-column label="指标名称" prop="kpi.name" width="120"></el-table-column>
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
      <el-table-column label="指标分数" width="90">
        <template slot-scope="scope">
            <el-input v-model="scope.row.score" clearable placeholder="请输入"></el-input>
        </template>
      </el-table-column> 
      <el-table-column label="评分人" width="230">
        <template slot-scope="scope">
          <span v-for="(item,index) in scope.row.PRIUs"
          :key="index">{{item.user.nickName}},</span>
        </template>
      </el-table-column>
        <el-table-column label="按钮组">
          <template slot-scope="scope">
            <el-button @click="performanceReviewItemDataEnter(scope.row)" type="primary" size="mini" slot="reference" label="修改">修改</el-button>
            <el-popover placement="top" width="160" v-model="scope.row.visible">
              <p>确定要删除吗？</p>
              <div style="text-align: right; margin: 0">
                <el-button size="mini" type="text" @click="scope.row.visible = false">取消</el-button>
                <el-button type="primary" size="mini" @click="deletePRI(scope.row)" :disabled="isDisable">确定</el-button>
              </div>
              <el-button type="danger" icon="el-icon-delete" size="mini" slot="reference">删除</el-button>
              <el-button @click="openRatioDialog(scope.row)" type="primary" size="mini" slot="reference" label="编辑人员权重">编辑人员权重</el-button>
            </el-popover>
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
    <!-- 添加指标 -->
    <el-dialog :before-close="closeKpiDialog" :visible.sync="kpiDialog" title="添加指标" width= "90%" :append-to-body="true" :fullscreen ="true">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="指标标签">
          <el-input placeholder="搜索条件" v-model="searchInfo.tagName"></el-input>
        </el-form-item> 
        <el-form-item label="指标名称">
          <el-input placeholder="搜索条件" v-model="searchInfo.name"></el-input>
        </el-form-item>    
        <el-form-item label="指标说明">
          <el-input placeholder="搜索条件" v-model="searchInfo.description"></el-input>
        </el-form-item>    
        <el-form-item>
          <el-button @click="searchKpiInfo" type="primary">查询</el-button>
        </el-form-item>
      </el-form>
    <el-table
      :data="kpiList"
      @selection-change="handleSelectionChange"
      border
      ref="multipleTable"
      stripe
      style="width: 100%"
      tooltip-effect="dark"
    >
    <el-table-column label="指标类型" width="120">
      <template slot-scope="scope">
        <span v-for="(item,index) in scope.row.Tags"
        :key="index">{{item.name}}<br/></span>
      </template>
    </el-table-column>
    <el-table-column label="指标名称" prop="name" width="120"></el-table-column> 
    
    <el-table-column label="指标说明" prop="description" width="460" type="textarea"></el-table-column> 
    
    <!-- <el-table-column label="指标状态" prop="Status" width="120"></el-table-column>  -->
    
    <el-table-column label="指标算法" prop="category" width="460" type="textarea"></el-table-column> 
     <el-table-column label="设置指标分数" width="160">
      <template slot-scope="scope">
          <el-input v-model="scope.row.evaluationKpis.kpiScore" clearable placeholder="请输入"></el-input>
      </template>
    </el-table-column>
    <el-table-column label="设置评分人" width="330">
      <template slot-scope="scope">
          <el-cascader
            @change="(val)=>{handleOptionChange(val,scope.row)}"
            v-model="scope.row.userId"
            :options="userOptions"
            :rules="rules"
            clearable
            :props="{ checkStrictly: true,label:'nickName',value:'id',multiple: true}"
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
    <el-pagination
      background
      :current-page="kpiPage"
      :page-size="kpiPageSize"
      :page-sizes="[5,10, 30, 50, 100]"
      :style="{float:'right',padding:'20px'}"
      :total="kpiTotal"
      @current-change="kpiStoreCurrentChange"
      @size-change="kpiStoreSizeChange"
      layout="total, sizes, prev, pager, next, jumper"
    ></el-pagination>
    </el-dialog>
    <el-dialog :before-close="closeRatioDialog" :visible.sync="ratioDialog" title="编辑人员权重" :append-to-body="true" style="width: 50%,marigin:right" :fullscreen ="true"
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
    <el-pagination
      background
      :current-page="ekuPage"
      :page-size="ekuPageSize"
      :page-sizes="[5,10,15,20]"
      :style="{float:'right',padding:'20px'}"
      :total="ekuTotal"
      @current-change="ekuCurrentChange"
      @size-change="ekuSizeChange"
      layout="total, sizes, prev, pager, next, jumper"
    ></el-pagination>
    </el-dialog>
    
    <el-dialog :before-close="closeUserDialog" :visible.sync="userDialog" title="增加评分人" :append-to-body="true" style="width: 50%,marigin:right" :fullscreen ="true"
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
    <el-pagination
      background
      :current-page="userPage"
      :page-size="userPageSize"
      :page-sizes="[5,10,15,20]"
      :style="{float:'right',padding:'20px'}"
      :total="userTotal"
      @current-change="userCurrentChange"
      @size-change="userSizeChange"
      layout="total, sizes, prev, pager, next, jumper"
    ></el-pagination>
    </el-dialog>
  </div>
</template>

<script>
import {
    createPerformanceReview,
    deletePerformanceReview,
    deletePerformanceReviewByIds,
    findPerformanceReview,
    getPRListWithoutFinishedStatus,
    updatePerformanceReviewByInfo,
    getLastPerformanceReview,
    updatePRStatysByIds,
    getPRByID
} from "@/api/pas/performanceReview";
import {
    getLastPRICreatePRIU,
    removePRIU,
    getPRIUByPRIID,
    updatePRIU,
    createPRIU
} from "@/api/pas/performanceReviewItemUser";
import {
    getEvaluationList,
    findEvaluation
} from "@/api/pas/evaluation";
import {
    getUserList,
    getUserListByAuthorityId
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
    updatePRItemStatysByIds,
    deletePRItemById,
    getLastPRI
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
      listApi: getPRListWithoutFinishedStatus,
      dialogFormVisible: false,
      kpiDialog:false,
      prItemDialog: false,
      userDialog: false,
      ratioDialog: false,
      visible: false,
      isDisable:false,
      dialogFlag: false,
      loading:false,
      dialogTitle:"新增指标",
      saveID:0,
      page: 1,
      pageSize: 10,
      ekuPage:1,
      ekuPageSize:10,
      userPage:1,
      userPageSize:10,
      userTotal:10,
      kpiPage: 1,
      priPage: 1,
      total: 10,
      kpiPageSize: 10,
      priPageSize: 10,
      type: "",
      searchInfo: {},
      priTotal:0,
      kpiTotal:0,
      ekuTotal:0,
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
      userData:{},
      rules: {
        name:[ { required: true, message: '请输入', trigger: 'blur' }],
        status:[ { required: true, message: '请输入', trigger: 'blur' }],
        description:[ { required: true, message: '请输入', trigger: 'blur' }],
        evaluationId:[ { required: true, message: '请输入', trigger: 'blur' }],
        employeeId:[ { required: true, message: '请输入', trigger: 'blur' }],
        startDate:[ { required: true, message: '请输入', trigger: 'blur' }],
        endingDate:[ { required: true, message: '请输入', trigger: 'blur' }],
      }
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
      //考核增加指标
      async kpiDataEnter(row){
        this.isDisable=true;
        if (row.userId==null) {
          this.$message({ type: 'warning', message: "请选择评分人" })
          this.isDisable=false;
        }
          var item = []
            item.push({
              score:Number(row.evaluationKpis.kpiScore),
              kpiId:row.ID,
              PRId:this.formData.ID,
              status:100,
              })
        const pr = await findPerformanceReview({ID:this.formData.ID})
        this.formData = pr.data.rePerformanceReview;
        const res = await createPerformanceReviewItem({item})
        if(res.code == 0){
          const lastPR = await getLastPRI()
          this.lastPRData = lastPR.data.rePRI
          var items = []
          for (let i = 0; i < row.userId.length; i++) {
            items.push({
              userId:Number(row.userId[i]),
              priid:Number(this.lastPRData.ID),
              score:Number((Number(1)/Number(row.userId.length)).toFixed(5)),   
            })
          }
            const cp = await createPRIU({
              items:items
            })
            if (cp.code == 0) {
              this.isDisable=false;
              this.$message({ type: 'success', message: "添加成功" })
              const sum =  Number(this.formData.score) + Number(row.evaluationKpis.kpiScore)
              updatePerformanceReviewByInfo({...this.formData,score:Number(sum)})
              const res = await getPerformanceReviewItemListById({ PRId: this.formData.ID,
                page: this.priPage, 
                pageSize: this.priPageSize})
              this.priTotal=res.data.total
              this.performanceReviewItemData = res.data.list;
            }
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
      async searchKpiInfo() {
        this.kpiPage = 1
        this.kpiPageSize = 10
        const ref = await getKpiList({...this.searchInfo,
          page: this.kpiPage, 
          pageSize: this.kpiPageSize})
        this.kpiTotal=ref.data.total
        this.kpiList=ref.data.list
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
      this.dialogTitle = "编辑考核表";
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
        const res = await getUserListByAuthorityId({
          page: this.userPage, 
          pageSize: this.userPageSize,
          authorityIds:["10000","8888"]})
        this.userTotal= res.data.total
        this.userData = res.data.list
        this.userDialog = true
    },
    async removeUser(row){
      this.isDisable=true;
      row.visible = false;
      const ref = await removePRIU({ID:row.id,priid:row.priid})
      if(ref.code ==0){
        this.$message({
          type: "success",
          message: "移除成功"
        });
        this.isDisable=false;
        const res = await getPRIUByPRIID({
          page: this.ekuPage, 
          pageSize: this.ekuPageSize,
          priid:row.priid})
        this.ekuTotal = res.data.total
        this.priuData = res.data.list
        const re = await getPerformanceReviewItemListById({ PRId: row.performanceReviewItem.PRId,
          page: this.priPage, 
          pageSize: this.priPageSize})
        this.priTotal=re.data.total
        this.performanceReviewItemData = re.data.list;
      }
    },
    async openRatioDialog(row) {
        const res = await getPRIUByPRIID({
          page: this.ekuPage, 
          pageSize: this.ekuPageSize,
          priid:row.ID})
        this.ekuTotal = res.data.total
        this.priuData = res.data.list
        this.saveData.priid=row.ID
        this.ratioDialog = true
      },
    async changeRatioData(row){
      this.isDisable=true;
      // if(row.score<=0||row.score>1){
      //   this.$message({
      //     type:"error",
      //     message:"无效输入"
      //   })
      //     this.isDisable=false;
      //     return
      //   }
      const ref = await updatePRIU({...row,
        score:Number(row.score),
        status:Number(row.status),
        result:Number(row.result),
        userId:Number(row.userId),
      })
        if (ref.code == 0) {
          this.$message({
            type:"success",
            message:"添加成功"
          })
          this.isDisable=false;
          const res = await getPRIUByPRIID({
            page: this.ekuPage, 
            pageSize: this.ekuPageSize,
            priid:row.priid})
          this.ekuTotal = res.data.total
          this.priuData = res.data.list
          const re = await getPerformanceReviewItemListById({ PRId: this.priuData[0].performanceReviewItem.PRId,
            page: this.priPage, 
            pageSize: this.priPageSize})
          this.priTotal=re.data.total
          this.performanceReviewItemData = re.data.list;
        }
    },
    async userDataEnter(row){
      var items=[{
        priid:this.saveData.priid,
        userid:row.ID,
        //score:Number((Number(1)/Number(this.priuData.length+1)).toFixed(5)),
        status:100,
      }]
      const ref = await createPRIU({
        items:items
      })
      if(ref.code == 0){
        const res = await getPRIUByPRIID({
          page: this.ekuPage, 
          pageSize: this.ekuPageSize,
          priid:this.saveData.priid})
        this.ekuTotal = res.data.total
        this.priuData = res.data.list
        const re = await getPerformanceReviewItemListById({ PRId: this.priuData[0].performanceReviewItem.PRId,
          page: this.priPage, 
          pageSize: this.priPageSize})
        this.priTotal=re.data.total
        this.performanceReviewItemData = re.data.list;
        this.$message({
          type:"success",
          message:"添加成功"
        })
      }
    },
    closeRatioDialog() {
      this.ratioDialog = false;
    },
    closeUserDialog() {
      this.userDialog = false;
    },
    // 初始化表单
    initForm() {
      if (this.$refs.prForm) {
        this.$refs.prForm.resetFields();
      }
      this.formData = {
            name:"",
            startDate:"",
            endingDate:"",
            status:"",
            evaluationId:"",
            employeeId:"",
            Tags: [],
      };
    },
    closeDialog() {
      this.initForm();
      this.dialogFormVisible = false;
      this.isDisable=false
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
    async ekuSizeChange(val) {
        this.ekuPageSize = val
        const res = await getPRIUByPRIID({
          page: this.ekuPage, 
          pageSize: this.ekuPageSize,
          priid:this.saveData.priid})
        this.ekuTotal = res.data.total
        this.priuData = res.data.list
    },
    async ekuCurrentChange(val) {
        this.ekuPage = val
        const res = await getPRIUByPRIID({
          page: this.ekuPage, 
          pageSize: this.ekuPageSize,
          priid:this.saveData.priid})
        this.ekuTotal = res.data.total
        this.priuData = res.data.list
    },
    async userSizeChange(val) {
        this.userPageSize = val
        const res = await getUserListByAuthorityId({
          page: this.userPage, 
          pageSize: this.userPageSize,
          authorityIds:["10000","8888"]})
        this.userTotal= res.data.total
        this.userData = res.data.list
    },
    async userCurrentChange(val) {
        this.userPage = val
        const res = await getUserListByAuthorityId({
          page: this.userPage, 
          pageSize: this.userPageSize,
          authorityIds:["10000","8888"]})
        this.userTotal= res.data.total
        this.userData = res.data.list
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
    async performanceReviewItemDataEnter(row){
        const res = await updatePerformanceReviewItemByInfo({
        id:Number(row.ID),
        score:Number(row.score),
        status:Number(row.status),
        });
        const prlist = await getPerformanceReviewItemListById({ PRId: this.formData.ID,
          page: this.priPage, 
          pageSize: this.priPageSize})
        this.priTotal=prlist.data.total
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
      this.$refs.prForm.validate(async valid => {
        if (valid) {
          this.loading=true
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
                  //userId:Number(this.evaluationKpiData[i].userId),
                  PRId:Number(this.formData.ID),
                  status:100,
                })
              }
            var re =await createPerformanceReviewItem({item})
            if (re.code == 0) {
              getLastPRICreatePRIU({
                ekuid:Number(this.formData.evaluationId),
                prid:Number(this.formData.ID)
              })
            }
              break;
            case "update":
              res = await updatePerformanceReviewByInfo({...this.formData,
              evaluationId:Number(this.formData.evaluationId),
              status:Number(this.formData.status),
              employeeId:Number(this.formData.employeeId)});
              this.isDisable=false
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
            this.loading=false
            this.getTableData()
            this.closeDialog();
          }
          this.initForm();
        }
      });
    },
    //删除考核表中的指标
    async deletePRI(row){
      this.isDisable = true;
      const res = await deletePRItemById({ ID: row.ID });
      if (res.code == 0) {
        row.visible = false;
        this.$message({
          type: "success",
          message: "删除成功"
        });
      const sum = this.formData.score-row.score
      updatePerformanceReviewByInfo({...this.formData,score:Number(sum)})
      const res = await getPerformanceReviewItemListById({ PRId: row.PRId,
        page: this.priPage, 
        pageSize: this.priPageSize})
      this.priTotal=res.data.total
      this.performanceReviewItemData = res.data.list;
      this.getTableData()
      }
      this.isDisable = false;
    },
    async onConfirm(){
      const ids = []
        if(this.multipleSelection.length == 0){
          this.$message({
            type: 'warning',
            message: '请选择要开始确认的考核'
          })
          this.confirmVisible = false
          return
        }
        for (let i = 0; i < this.multipleSelection.length; i++) {
            if (this.multipleSelection[i].status!=100) {
              this.$message({
                type: 'warning',
                message: '考核已经开始'
              })
              this.confirmVisible = false
              return
            }
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
          this.confirmVisible = false
          updatePRItemStatysByIds({ids:ids,status:1})
          this.getTableData()
        }
    },
    async openDialog() {
      this.dialogTitle = "创建考核表";
      this.initForm();
      this.formData.status=100
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
    const evaluations = await getEvaluationList({ status:1,page: 1, pageSize: 999 });
    this.setEvaluationOptions(evaluations.data.list);
}
};
</script>

<style>
</style>