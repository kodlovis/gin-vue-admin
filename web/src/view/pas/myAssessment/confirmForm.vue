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
      <h1>当前考核概览</h1>
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
        <el-table-column label="权重分值" prop="score" width="120"></el-table-column> 
        <el-table-column label="开始时间" prop="startDate" width="120"></el-table-column>
        <el-table-column label="结束时间" prop="endingDate" width="120"></el-table-column>
        </el-table>
      <h5>当前考核表详情</h5>
      <el-table
          :data="prData"
          @selection-change="handleSelectionChange"
          border
          ref="multipleTable"
          stripe
          style="width: 100%"
          tooltip-effect="dark"
        >
        <el-table-column label="考核名称" prop="prs.name" width="120"></el-table-column> 
        <el-table-column label="指标名称" prop="kpi.name" width="120"></el-table-column> 
        <el-table-column label="指标算法" prop="kpi.category" width="460"></el-table-column> 
        <el-table-column label="指标描述" prop="kpi.description" width="460"></el-table-column> 
        <el-table-column label="评分人" prop="user.nickName" width="120">
          <template slot-scope="scope">
            <span v-for="(item,index) in scope.row.PRIUs"
            :key="index">{{item.user.nickName}}<br/></span>
          </template>
        </el-table-column> 
        <el-table-column label="权重分值" prop="score" width="120"></el-table-column>
          <el-table-column label="按钮组">
            <template slot-scope="scope">
              <el-button class="table-button" @click="confirmKpi(scope.row)" size="small" type="primary" icon="el-icon-edit" :disabled="isDisable">确认</el-button>
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
    

  </div>
</template>

<script>
import {
    updatePRI,
    getPRItemListByStatusPrid,
} from "@/api/pas/performanceReviewItem";  //  此处请自行替换地址
import {    
    getPRListByUser,
} from "@/api/pas/performanceReview";  //  此处请自行替换地址
import {    
    updatePRIUStatusByPRIID
} from "@/api/pas/performanceReviewItemUser";  //  此处请自行替换地址

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
      isDisable:false,
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
        prs:{
          name:"",
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
      this.isDisable=true
      const res = await updatePRIUStatusByPRIID({
          ID:row.ID,
          status:92,
          prStatus:3,
      })
      if(res.code == 0){
          this.$message({
          type: "success",
          message: "确认成功"})
          this.getPRListByUser()
      }
    },
    async rejectKpi(row){
      const res = await updatePRI({...row,
          status:99,
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
          ids:[2],
          })
      this.acData = res.data.list
      const num = [] 
      for (let i = 0; i < this.acData.length; i++) {
        num.push(this.acData[i].ID)
      }
      const prItem = await getPRItemListByStatusPrid({
          ids: num,
          status:2,
          })
      this.prData = prItem.data.list 
      this.isDisable=false
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