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
        <el-form-item>
          <el-button @click="searchPRIInfo" type="primary">查询</el-button>
        </el-form-item>
      </el-form>
    </div>
      <el-table
          :data="acData"
          @selection-change="handleSelectionChange"
          border
          ref="multipleTable"
          stripe
          style="width: 100%"
          tooltip-effect="dark"
        >
        
        <el-table-column label="考核名称" prop="name" width="240"></el-table-column> 
        
        <el-table-column label="考核状态" prop="status" width="120">
            <template slot-scope="scope">
                <!-- <span>{{scope.row.status==0?"评分人确认":""}}</span> -->
                <span>{{filterDict(scope.row.status)}}</span>
            </template>
         </el-table-column> 
        <el-table-column label="被考核人" prop="user.nickName" width="120"></el-table-column> 
        <el-table-column label="总分" prop="score" width="120"></el-table-column> 
        <el-table-column label="得分" prop="result" width="120"></el-table-column> 
        <el-table-column label="开始时间" prop="startDate" width="120">
          <template slot-scope="scope">{{scope.row.startDate|formatDate}}</template>
        </el-table-column>
        <el-table-column label="结束时间" prop="endingDate" width="120">
          <template slot-scope="scope">{{scope.row.endingDate|formatDate}}</template>
        </el-table-column>
        <el-table-column label="备注" prop="comment" width="400">
        </el-table-column>
        
          <el-table-column label="按钮组">
            <template slot-scope="scope">
                <el-button class="table-button" @click="getPRI(scope.row)" size="small" type="primary" >查看指标</el-button>
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
        sortable
      >
      <el-table-column label="指标类型" width="120" sortable>
        <template slot-scope="scope">
            <span v-for="(item,index) in scope.row.kpi.Tags"
        :key="index">{{item.name}}</span>
        </template>
      </el-table-column>
      <el-table-column label="指标名称" prop="kpi.name" width="150"></el-table-column>

      <el-table-column label="指标说明" prop="kpi.description" width="360" type="textarea"></el-table-column>
      
      <el-table-column label="指标算法" prop="kpi.category" width="360" type="textarea"> </el-table-column> 
      <el-table-column label="指标状态" prop="status" width="160">
      <template slot-scope="scope">
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
        :page-sizes="[1,5,10, 30, 50, 100]"
        :style="{float:'right',padding:'20px'}"
        :total="priTotal"
        @current-change="kpiCurrentChange"
        @size-change="kpiSizeChange"
        layout="total, sizes, prev, pager, next, jumper"
      ></el-pagination>
      </el-dialog>
    </div>
  </div>
</template>

<script>
import {
    getPerformanceReviewItemListById
} from "@/api/pas/performanceReviewItem";  //  此处请自行替换地址
import {    
    getPRListByUser,
    getPRByID
} from "@/api/pas/performanceReview";  //  此处请自行替换地址

import { formatTimeToStr } from "@/utils/date";
import { getDict } from "@/utils/dictionary";
import infoList from "@/mixins/infoList";
import { mapGetters } from "vuex";
export default {
  name: "result",
  mixins: [infoList],
  data() {
    return {
      listApi: getPRListByUser,
      type: "",
      dictList:[],
      multipleSelection: [],
      countData:9,
      saveID:0,
      page: 1,
      pageSize: 10,
      total: 10,
      priPage:1,
      priPageSize:5,
      priTotal:0,
      prItemDialog: false,
      kpiDictList:[],
      searchInfo: {},
      acData:{
        startDate:new Date(),
        endingDate:new Date(),
        result:0,
        score:0,
        status:0,
        comment:"",
        name:"",
        user:{
            nickName:"",
        },
      },
        performanceReviewItemData:{
        },
    };
  },
  computed: {
    ...mapGetters("user", ["userInfo", "token"])
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
    //打开查看指标的弹窗并传参
    async getPRI(row) {
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
        //获取指标状态字典
        const kpi = await getDict("kpi");
        kpi.map(item=>item.value)
        this.kpiDictList = kpi
        this.prItemDialog = true;
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
      //条件搜索前端看此方法
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
      handleSelectionChange(val) {
        this.multipleSelection = val
      },
    async searchPRIInfo() {
      this.page = 1
      this.pageSize = 10
      const ref = await getPRListByUser({...this.searchInfo,
          ID:this.userInfo.ID,
          page: this.page, 
          pageSize: this.pageSize})
      this.total=ref.data.total
      this.acData=ref.data.list
    },
    async getPRListByUser(){
      const res = await getPRListByUser({
          ID:this.userInfo.ID,
          page: this.page, 
          pageSize: this.pageSize
          })
      this.acData = res.data.list
    },
    async handleSizeChange(val) {
        this.pageSize = val
        const ref = await getPRListByUser({...this.searchInfo,
            ID:this.userInfo.ID,
            page: this.page, 
            pageSize: this.pageSize})
        this.total=ref.data.total
        this.acData=ref.data.list
    },
    async handleCurrentChange(val) {
        this.page = val
        const ref = await getPRListByUser({...this.searchInfo,
            ID:this.userInfo.ID,
            page: this.page, 
            pageSize: this.pageSize})
        this.total=ref.data.total
        this.acData=ref.data.list
    },
  },
  async created() {
    //获取考核状态字典
    const pr = await getDict("PR");
    pr.map(item=>item.value)
    this.dictList = pr
    this.getPRListByUser()
    
}
};
</script>

<style>
</style>