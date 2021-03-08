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
      <h1 style="font-size:20px">当前考核概览</h1>
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
        <el-table-column label="考核总分" prop="score" width="120"></el-table-column> 
        <el-table-column label="考核得分" prop="result" width="120"></el-table-column> 
        <el-table-column label="开始时间" prop="startDate" width="120">
          <template slot-scope="scope">{{scope.row.startDate|formatDate}}</template>
        </el-table-column>
        <el-table-column label="结束时间" prop="endingDate" width="120">
          <template slot-scope="scope">{{scope.row.endingDate|formatDate}}</template>
        </el-table-column>
        <el-table-column label="备注" prop="score" width="520">
          <template slot-scope="scope">
            <el-input v-model="scope.row.comment" clearable placeholder="请输入备注"></el-input>
          </template>
        </el-table-column>
        
          <el-table-column label="按钮组">
            <template slot-scope="scope">
              <el-button class="table-button" @click="confirmKpi(scope.row)" size="small" type="primary" icon="el-icon-edit">确认</el-button>
              <el-popover placement="top" width="160" v-model="scope.row.visible">
                <p>确定要驳回吗？</p>
                <div style="text-align: right; margin: 0">
                  <el-button size="mini" type="text" @click="scope.row.visible = false">取消</el-button>
                  <el-button type="primary" size="mini" @click="rejectKpi(scope.row)">确定</el-button>
                </div>
                <el-button type="danger" icon="el-icon-delete" size="mini" slot="reference">驳回</el-button>
              </el-popover>
            </template>
          </el-table-column>
        </el-table>
      <h5 style="font-size:20px">当前考核表详情</h5>
      <el-table
          :data="prData"
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
            :key="index">{{item.name}}<br/></span>
          </template>
        </el-table-column>
        <el-table-column label="指标名称" prop="kpi.name" width="120"></el-table-column> 
        
        <el-table-column label="指标算法" prop="kpi.category" width="460"></el-table-column> 
        <el-table-column label="指标描述" prop="kpi.description" width="460"></el-table-column> 
        <el-table-column label="指标总分" prop="score" width="100"></el-table-column>
        <el-table-column label="指标得分" prop="result" width="100"></el-table-column>
        <el-table-column label="评分人及评分" prop="PRIUs.user.nickName" width="160">
          <template slot-scope="scope">
            <span v-for="(item,index) in scope.row.PRIUs"
            :key="index">{{item.user.nickName}}:{{item.result}}分<br/></span>
          </template>
        </el-table-column> 
        </el-table>
    

  </div>
</template>

<script>
import {
    getPRItemCount,
    getPRItemListByPrids
} from "@/api/pas/performanceReviewItem";  //  此处请自行替换地址
import {    
    updatePRStatusById,
    getPRListByUser,
    findPerformanceReview
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
      prData:{
        result:"",
        kpi:{
            name:"",
            category:"",
            description:"",
            Tags:{
                name:"",
            }
        },
      },
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
        }
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
    async confirmKpi(row) {
        //查询已进行反馈的指标数量
        const count = await getPRItemCount({
            PRId:row.ID,
            status: 7,
        })
        this.countData = count.data.count;
        //如果都已经反馈除以总数=1，完成本次考核
        if(this.countData == 1){
          const res = await updatePRStatusById({
                ID:row.ID,
                status:9,
                result:row.result
          })
          if(res.code == 0) {
            this.getPRListByUser()
            this.$message({
            type: "success",
            message: "确认成功"})
            return
          }
        }
        //检查考核表状态，防止评分人的操作而导致的冲突录入
        const pr = await findPerformanceReview({ID:row.ID})
        if(pr.data.rePerformanceReview.status==9){
          this.getPRListByUser()
          this.$message({
          type: "success",
          message: "确认成功"})
          return
        }
        //更改考核状态为已确认
        const res = await updatePRStatusById({
          ID:row.ID,
          status:8,
          result:row.result
        }) 
        if (res.code == 0) {
          this.getPRListByUser()
          this.$message({
          type: "success",
          message: "确认成功"})
        }
    },
    async rejectKpi(row){
      const res = await updatePRStatusById({
          ID:row.ID,
          status:999,
          result:row.result
      })
      if(res.code == 0){
          this.getPRListByUser()
          this.$message({
          type: "success",
          message: "驳回成功"})
      }
    },
    async getPRListByUser(){
      const res = await getPRListByUser({
          ID:this.userInfo.ID,
          ids:[7,8,999],
          })
      this.acData = res.data.list
      const num = [] 
      for (let i = 0; i < this.acData.length; i++) {
        num.push(this.acData[i].ID)
      }
      const prItem = await getPRItemListByPrids({ids:num})
      this.prData = prItem.data.list 
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