<template>
  
  <div v-loading="loading">
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
      <el-table
          :data="acData"
          @selection-change="handleSelectionChange"
          border
          ref="multipleTable"
          stripe
          style="width: 100%"
          tooltip-effect="dark"
        >
        
        <el-table-column label="指标名称" prop="kpi.name" width="120"></el-table-column> 
        
        <el-table-column label="指标算法" prop="kpi.category" width="460"></el-table-column> 
        <el-table-column label="指标描述" prop="kpi.description" width="460"></el-table-column> 
        <el-table-column label="被考评人" prop="prs.user.nickName" width="120"></el-table-column> 
        <el-table-column label="权重分值" prop="score" width="120"></el-table-column>
        <el-table-column label="得分">
          <template slot-scope="scope">
            <el-input v-model="scope.row.result" clearable placeholder="请输入分数"></el-input>
          </template>
        </el-table-column>
          <el-table-column label="按钮组">
            <template slot-scope="scope">
              <el-button class="table-button" @click="confirmKpi(scope.row)" size="small" type="primary" icon="el-icon-edit" >确认</el-button>
              <!-- <el-popover placement="top" width="160" v-model="scope.row.visible">
                <p>确定要驳回吗？</p>
                <div style="text-align: right; margin: 0">
                  <el-button size="mini" type="text" @click="scope.row.visible = false">取消</el-button>
                  <el-button type="primary" size="mini" @click="rejectKpi(scope.row)">确定</el-button>
                </div>
                <el-button type="danger" icon="el-icon-delete" size="mini" slot="reference">驳回</el-button>
              </el-popover> -->
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

  </div>
</template>

<script>
import {
    getPRItemListByUser,
    updatePRItemStatusById,
    getPRItemCount,
    updatePRItemStatusByPrId
} from "@/api/pas/performanceReviewItem";  //  此处请自行替换地址
import {    
    updatePRStatusById,
    findPerformanceReview
} from "@/api/pas/performanceReview";  //  此处请自行替换地址

import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";
import { mapGetters } from "vuex";
export default {
  name: "assessmentForm",
  mixins: [infoList],
  data() {
    return {
      listApi: getPRItemListByUser,
      type: "",
      multipleSelection: [],formData:[],
      countData:9,
      loading:false,
      page: 1,
      total: 10,
      pageSize: 10,
      acData:{
        score:0,
        result:0,
        kpi:{
          name:"",
          category:"",
          description:"",
        },
        prs:{
          user:{
            nickName:"",
          },
        },
        user:{
          nickName:"",
        },
      },
      prData:{
        result:0,
      }
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
      handleSelectionChange(val) {
        this.multipleSelection = val
      },
    async confirmKpi(row) {
      this.loading=true
      if(row.result>row.score||row.result<0){
          this.$message({
            type: 'warning',
            message: '无效输入'
          })
          this.loading=false
          return
      }
      const res = await updatePRItemStatusById({
          ID:row.ID,
          status:5,
          result:Number(row.result),
      })
      if(res.code == 0){
          this.getPRItemListByUser()
          this.$message({
          type: "success",
          message: "确认成功"})
        const pr = await findPerformanceReview({ID:row.PRId})
        this.prData = pr.data.rePerformanceReview
        const total=Number(row.result)+Number(this.prData.result)
        const update = await updatePRStatusById({
          ID:row.PRId,
          result:Number(total),
          status:5,
        })
        if(update.code==0){
          const count = await getPRItemCount({
              PRId:row.PRId,
              status: 5,
          })
          this.countData = count.data.count;
          if(this.countData == 1){
              updatePRStatusById({
                  ID:row.PRId,
                  result:Number(total),
                  status: 7,
              })
              updatePRItemStatusByPrId({
                PRId:row.PRId,
                status:6,
            })
          }
        }
      }
          this.loading=false
    },
    handleSizeChange(val) {
        this.pageSize = val
        this.getPRItemListByUser()
    },
    handleCurrentChange(val) {
        this.page = val
        this.getPRItemListByUser()
    },
    async getPRItemListByUser(page = this.page, pageSize = this.pageSize){
      const res = await getPRItemListByUser({
          ID:this.userInfo.ID,
          status:4,
          page: page, 
          pageSize: pageSize
          })
      this.total = res.data.total
      this.page = res.data.page
      this.pageSize = res.data.pageSize
      this.acData = res.data.list
    },
  },
  async created() {
    await this.getTableData();
    this.getPRItemListByUser()
}
};
</script>

<style>
</style>