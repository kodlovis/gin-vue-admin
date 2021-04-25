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
          :rules="rules" 
        >
        <el-table-column label="考核名称" prop="performanceReviewItem.prs.name" width="120"></el-table-column> 
        <el-table-column label="指标类型" width="120">
          <template slot-scope="scope">
            <span v-for="(item,index) in scope.row.performanceReviewItem.kpi.Tags"
            :key="index">{{item.name}}<br/></span>
          </template>
        </el-table-column>
        <el-table-column label="指标名称" prop="performanceReviewItem.kpi.name" width="120"></el-table-column> 
        <el-table-column label="指标算法" prop="performanceReviewItem.kpi.category" width="400"></el-table-column> 
        <el-table-column label="指标描述" prop="performanceReviewItem.kpi.description" width="400"></el-table-column> 
        <el-table-column label="被考评人" prop="performanceReviewItem.prs.user.nickName" width="100"></el-table-column> 
        <!-- <el-table-column label="权重分值" width="120">
          <template slot-scope="scope">
            <span>{{scope.row.score*100+"%"}}<br/></span>
          </template>
          </el-table-column> -->
        <el-table-column label="分值" prop="performanceReviewItem.score" width="100"></el-table-column>
        <el-table-column label="得分" width="170">
          <template slot-scope="scope">
            <el-input 
              size="small" 
              v-model="scope.row.result" 
              placeholder="请输入分数"  
              :min="0"
              type="number"
              onKeypress="return (/[\d.]/.test(String.fromCharCode(event.keyCode)))"
            ></el-input>
          </template>
        </el-table-column>
          <el-table-column label="按钮组">
            <template slot-scope="scope">
              <el-button class="table-button" @click="confirmKpi(scope.row)" size="small" type="primary" icon="el-icon-edit"  :disabled="isDisable">确认</el-button>
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
    updatePRItemStatusById,
    getPRItemCount,
    updatePRItemStatusByPrId,
} from "@/api/pas/performanceReviewItem";  //  此处请自行替换地址
import {    
    updatePRStatusById,
    updatePRResult
} from "@/api/pas/performanceReview";
import {    
    getPRIUListByUser
} from "@/api/pas/performanceReviewItemUser";


import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";
import { mapGetters } from "vuex";
export default {
  name: "assessmentForm",
  mixins: [infoList],
  data() {
    return {
      listApi: getPRIUListByUser,
      type: "",
      multipleSelection: [],formData:[],
      countData:9,
      loading:false,
      isDisable:false,
      page: 1,
      total: 10,
      pageSize: 10,
      rules: {
        name:[ { required: true, message: '请输入', trigger: 'blur' }],
        status:[ { required: true, message: '请输入', trigger: 'blur' }],
        category:[ { required: true, message: '请输入', trigger: 'blur' }],
        description:[ { required: true, message: '请输入', trigger: 'blur' }],
        result:[ { required: true, message: '请输入', trigger: 'blur' }]
      },
      acData:{
        score:0,
        result:"",
        kpi:{
          name:"",
          category:"",
          description:"",
        },
      saveData:{
        result:"",
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
      this.isDisable=true
      if(row.result>row.performanceReviewItem.score||row.result<0||row.result==""){
          this.$message({
            type: 'warning',
            message: '无效输入'
          })
          this.loading=false
          this.isDisable=false
          return
      }
      const res = await updatePRItemStatusById({
          ID:row.id,
          status:5,
          result:Number(row.result),
      })
      if(res.code == 0){
          this.getPRIUListByUserNoRefresh()
          this.$message({
          type: "success",
          message: "确认成功"})
        const update = await updatePRResult({
          ID:row.performanceReviewItem.PRId,
        })
        if(update.code==0){
          const count = await getPRItemCount({
              PRId:row.performanceReviewItem.PRId,
              status: 5,
          })
          this.countData = count.data.count;
          if(this.countData == 1){
              updatePRStatusById({
                  ID:row.performanceReviewItem.PRId,
                  status: 7,
              })
              updatePRItemStatusByPrId({
                PRId:row.performanceReviewItem.PRId,
                status:6,
            })
          }
        }
      }
        this.isDisable=false
    },
    handleSizeChange(val) {
        this.pageSize = val
        this.getPRIUListByUser()
    },
    handleCurrentChange(val) {
        this.page = val
        this.getPRIUListByUser()
    },
    async getPRIUListByUserNoRefresh(page = this.page, pageSize = this.pageSize){
      const num = []
      for (let i = 0; i < this.acData.length; i++) {
        num.push(this.acData[i].result)
      }
      const res = await getPRIUListByUser({
          ID:this.userInfo.ID,
          status:4,
          page: page, 
          pageSize: pageSize
          })
      this.total = res.data.total
      this.page = res.data.page
      this.pageSize = res.data.pageSize
      this.acData = res.data.list
      for (let i = 0; i < this.acData.length; i++) {
        this.acData[i].result=num[i+1]
      }
    },
    async getPRIUListByUser(page = this.page, pageSize = this.pageSize){
      const res = await getPRIUListByUser({
          ID:this.userInfo.ID,
          status:4,
          page: page, 
          pageSize: pageSize
          })
      this.total = res.data.total
      this.page = res.data.page
      this.pageSize = res.data.pageSize
      this.acData = res.data.list
      for (let i = 0; i < this.acData.length; i++) {
        this.acData[i].result=""
      }
    },
  },
  async created() {
    this.getPRIUListByUser()
}
};

</script>

<style>
</style>