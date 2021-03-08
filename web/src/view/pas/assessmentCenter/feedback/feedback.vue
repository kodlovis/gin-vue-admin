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
    <div>
      <el-form>
        <el-form-item>
          <el-popover placement="top" v-model="skipVisible" width="160">
            <p>确定要跳过吗？</p>
              <div style="text-align: right; margin: 0">
                <el-button @click="skipVisible = false" size="mini" type="text">取消</el-button>
                <el-button @click="onSkip" size="mini" type="primary">确定</el-button>
              </div>
            <el-button size="mini" slot="reference" type="primary">批量跳过</el-button>
          </el-popover>
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
        <el-table-column type="selection" width="55"></el-table-column>
        <el-table-column label="考核名称" prop="performanceReviewItem.prs.name" width="120"></el-table-column> 
        <el-table-column label="指标类型" width="120">
          <template slot-scope="scope">
            <span v-for="(item,index) in scope.row.performanceReviewItem.kpi.Tags"
            :key="index">{{item.name}}<br/></span>
          </template>
        </el-table-column>
        <el-table-column label="指标名称" prop="performanceReviewItem.kpi.name" width="100"></el-table-column> 
        
        <el-table-column label="指标算法" prop="performanceReviewItem.kpi.category" width="380"></el-table-column> 
        <el-table-column label="指标描述" prop="performanceReviewItem.kpi.description" width="380"></el-table-column> 
        <el-table-column label="被考评人" prop="user.nickName" width="100"></el-table-column> 
        <el-table-column label="权重分值" prop="score" width="80"></el-table-column>
        <el-table-column label="反馈" width="200">
          <template slot-scope="scope">
            <el-input v-model="scope.row.comment" clearable placeholder="请输入反馈"></el-input>
          </template>
        </el-table-column>
          <el-table-column label="按钮组">
            <template slot-scope="scope">
              <el-button class="table-button" @click="confirmKpi(scope.row)" size="small" type="primary" icon="el-icon-edit" >提交</el-button>
              <el-button class="table-button" @click="skip(scope.row)" size="small" type="primary" icon="el-icon-edit" >跳过</el-button>
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

  </div>
</template>

<script>
import {
    getPRItemListByUser,
    updatePRItemStatusById,
    getPRItemCount
} from "@/api/pas/performanceReviewItem";  //  此处请自行替换地址
import {    
    updatePRStatusById,
    findPerformanceReview
} from "@/api/pas/performanceReview";  //  此处请自行替换地址

import {    
    getPRIUListByUser,
    updatePRIUStatustByIds
} from "@/api/pas/performanceReviewItemUser";
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
      page: 1,
      total: 10,
      pageSize: 10,
      loading:false,
      skipVisible:false,
      acData:{
        comment:"",
        score:0,
        result:0,
        kpi:{
          name:"",
          category:"",
          description:"",
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
      async onSkip(){
        const ids = []
        if(this.multipleSelection.length == 0){
          this.$message({
            type: 'warning',
            message: '请选择要跳过的数据'
          })
          return
        }
        this.multipleSelection &&
          this.multipleSelection.map(item => {
            ids.push(item.id)
          })
        const res = await updatePRIUStatustByIds({ ids:ids,status:7 })
        if (res.code == 0) {
          this.$message({
            type: 'success',
            message: '跳过成功'
          })
          this.skipVisible = false
          this.getPRIUListByUser()
        }
      },
      //条件搜索前端看此方法
      handleSelectionChange(val) {
        this.multipleSelection = val
      },
      async skip(row){
        this.loading=true
        //更新PRItem
        const res = await updatePRItemStatusById({
            ID:row.id,
            status:7,
            result:Number(row.result),
            comment:"无"
        })
      if(res.code == 0){
        //查询未反馈的数量
        const count = await getPRItemCount({
            PRId:row.performanceReviewItem.PRId,
            status: 6,
        })
        this.countData = count.data.total;
        //查询这条指标的考核表
        const pr = await findPerformanceReview({ID:row.performanceReviewItem.PRId})
        this.prData = pr.data.rePerformanceReview
        //如果未考核的数量为0
        if(this.countData == 0){
            //如果考核表的状态为已确认
            if (this.prData.status==8) {
                //将考核表设为已完成
                updatePRStatusById({
                    result:this.prData.result,
                    ID:row.performanceReviewItem.PRId,
                    status: 9,
                })
            }
        }
        //刷新列表
        this.getPRIUListByUser()
        this.$message({
        type: "success",
        message: "确认成功"})
        this.loading=false
      }
      },
    async confirmKpi(row) {
      this.loading=true
      //更新PRItem，写入反馈
      const res = await updatePRItemStatusById({
          ID:row.id,
          status:7,
          result:Number(row.result),
          comment:row.comment,
      })
      if(res.code == 0){
        //查询已反馈的数量
        const count = await getPRItemCount({
            PRId:row.performanceReviewItem.PRId,
            status: 7,
        })
        this.countData = count.data.count;
        //查询这条指标的考核表
        const pr = await findPerformanceReview({ID:row.performanceReviewItem.PRId})
        this.prData = pr.data.rePerformanceReview
        //如果已反馈除以总数=1
        if(this.countData == 1){
            //如果考核表的状态为已确认
            if (this.prData.status==8) {
                //将考核表设为已完成
                updatePRStatusById({
                    result:this.prData.result,
                    ID:row.performanceReviewItem.PRId,
                    status: 9,
                })
            }
        }
        //刷新列表
        this.getPRIUListByUser()
        this.$message({
        type: "success",
        message: "确认成功"})
        this.loading=false
      }
    },
    handleSizeChange(val) {
        this.pageSize = val
        this.getPRIUListByUser()
    },
    handleCurrentChange(val) {
        this.page = val
        this.getPRIUListByUser()
    },
    async getPRIUListByUser(page = this.page, pageSize = this.pageSize){
      const res = await getPRIUListByUser({
          ID:this.userInfo.ID,
          status:6,
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
    this.getPRIUListByUser()
}
};
</script>

<style>
</style>