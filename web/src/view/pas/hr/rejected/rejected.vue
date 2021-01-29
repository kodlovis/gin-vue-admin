<template>
  <div>
    <!-- <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="绩效考核表名称">
          <el-input placeholder="搜索条件" v-model="searchInfo.name"></el-input>
        <el-form-item label="考核表状态">
          <el-input placeholder="搜索条件" v-model="searchInfo.status"></el-input>
        </el-form-item>   
        </el-form-item>          
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
        </el-form-item>
      </el-form>
    </div> -->
      <h5>驳回的考核内容</h5>
      <el-table
          :data="acData"
          @selection-change="handleSelectionChange"
          border
          ref="multipleTable"
          stripe
          style="width: 100%"
          tooltip-effect="dark"
        >
        
        <el-table-column label="考核表名称" prop="prs.name" width="120"></el-table-column> 
        <el-table-column label="指标名称" prop="kpi.name" width="120"></el-table-column> 
        <el-table-column label="指标状态" prop="status" width="120">
            <template slot-scope="scope">
                <!-- <span>{{scope.row.status==0?"评分人确认":""}}</span> -->
                <span>{{filterKpiDict(scope.row.status)}}</span>
            </template>
        </el-table-column>
        <el-table-column label="指标算法" prop="kpi.category" width="460"></el-table-column> 
        <el-table-column label="指标描述" prop="kpi.description" width="460"></el-table-column> 
        <el-table-column label="被考评人" prop="prs.user.nickName" width="120"></el-table-column> 
        <el-table-column label="评分人" prop="user.nickName" width="120"></el-table-column> 
        <el-table-column label="权重分值" prop="score" width="120"></el-table-column>
        </el-table>
      <h5>驳回的考核结果</h5>
      <el-table
          :data="prData"
          @selection-change="handleSelectionChange"
          border
          ref="multipleTable"
          stripe
          style="width: 100%"
          tooltip-effect="dark"
        >
        <el-table-column label="考核表名称" prop="name" width="120"></el-table-column> 
        <el-table-column label="指标状态" prop="status" width="120">
            <template slot-scope="scope">
                <!-- <span>{{scope.row.status==0?"评分人确认":""}}</span> -->
                <span>{{filterDict(scope.row.status)}}</span>
            </template>
        </el-table-column>
        <el-table-column label="被考核人" prop="user.nickName" width="120"></el-table-column> 
        <el-table-column label="权重分值" prop="score" width="120"></el-table-column>
        <el-table-column label="得分" prop="result" width="120"></el-table-column>
        <el-table-column label="备注" prop="comment" width="360"></el-table-column>
          <el-table-column label="按钮组">
            <template slot-scope="scope">
              <el-button class="table-button" @click="getPRItem(scope.row)" size="small" type="primary" icon="el-icon-edit">查看</el-button>
            </template>
          </el-table-column>
        </el-table>
    <el-dialog :before-close="closeprItemDialog" :visible.sync="prItemDialog" title="编辑指标" :append-to-body="true" width="80%">
      <el-table
      :data="prItemData"
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
    <el-table-column label="指标状态" prop="kpi.status" width="90">
      <template slot-scope="scope">
          <!-- <span>{{scope.row.status==0?"评分人确认":""}}</span> -->
          <span>{{filterKpiDict(scope.row.status)}}</span>
      </template></el-table-column>
    <el-table-column label="权重分值" prop="score" width="90"></el-table-column> 
    <el-table-column label="得分" prop="result" width="90"></el-table-column> 
    <el-table-column label="评分人" prop="user.nickName" width="120"></el-table-column>
    <el-table-column label="考评人反馈" prop="comment" width="360"></el-table-column>
    </el-table>
    </el-dialog>
    <el-pagination
      :current-page="page"
      :page-size="pageSize"
      :page-sizes="[10]"
      :style="{float:'right',padding:'20px'}"
      :total="total"
      @current-change="handleCurrentChange"
      @size-change="handleSizeChange"
      layout="total, sizes, prev, pager, next, jumper"
    ></el-pagination>

  </div>
</template>

<script>
import {
    getPRItemListByStatus,
    getPerformanceReviewItemListById
} from "@/api/pas/performanceReviewItem";  //  此处请自行替换地址
import {    
    getPRBystatus
} from "@/api/pas/performanceReview";  //  此处请自行替换地址
import { getDict } from "@/utils/dictionary";
import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";
export default {
  name: "rejected",
  mixins: [infoList],
  data() {
    return {
      type: "",
      kpiDictList:[],
      multipleSelection: [],formData:[],
      countData:9,
      dictList:[],
      prItemDialog: false,
      page: 1,
      total: 10,
      pageSize: 10,
      prData:{
        name:"",
        comment:"",
        score:"",
        result:"",
        user:{
          nickName:"",
        }
      },
      acData:{
        score:0,
        kpi:{
          name:"",
          category:"",
          description:"",
        },
        user:{
          nickName:"",
        },
        prs:{
          name:"",
          user:{
            nickName:""
          }
        }
      },
      prItemData:{
        score:0,
        kpi:{
          name:"",
          category:"",
          description:"",
        },
        user:{
          nickName:"",
        },
        prs:{
          name:"",
          user:{
            nickName:""
          }
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
      //条件搜索前端看此方法
      handleSelectionChange(val) {
        this.multipleSelection = val
      },
    async getPRItem(row) {
      const res = await getPerformanceReviewItemListById({
          PRId:row.ID,
      })
      if (res.code == 0) {
        this.prItemData = res.data.list;
        this.prItemDialog = true;
      }
    },
    // async rejectKpi(row){
    //   const res = await updatePRItemStatusById({
    //       ID:row.ID,
    //       status:99,
    //   })
    //   if(res.code == 0){
    //       this.getPRItemListByStatus()
    //       this.$message({
    //       type: "success",
    //       message: "驳回成功"})
    //   }
    // },
    async getPRItemListByStatus(page = this.page, pageSize = this.pageSize){
      const res = await getPRItemListByStatus({
          status:99,
          page: page, 
          pageSize: pageSize
          })
      const pr = await getPRBystatus({
          status:99,
          page: page, 
          pageSize: pageSize
          })
      this.total = res.data.total
      this.page = res.data.page
      this.pageSize = res.data.pageSize
      this.acData = res.data.list
      this.prData = pr.data.list
    },
    handleSizeChange(val) {
        this.pageSize = val
        this.getPRItemListByStatus()
    },
    handleCurrentChange(val) {
        this.page = val
        this.getPRItemListByStatus()
    },
  },
  async created() {
    //获取指标状态字典
    const kpi = await getDict("kpi");
    kpi.map(item=>item.value)
    this.kpiDictList = kpi
    //获取考核状态字典
    const pr = await getDict("PR");
    pr.map(item=>item.value)
    this.dictList = pr
    this.getPRItemListByStatus()
}
};
</script>

<style>
</style>