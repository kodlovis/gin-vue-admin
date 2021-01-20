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
        
        <el-table-column label="考核表状态" prop="status" width="460">
          <template slot-scope="scope">
            <!-- <span>{{scope.row.status==0?"评分人确认":""}}</span> -->
            <span>{{filterDict(scope.row.status)}}</span>
          </template>
          </el-table-column> 
        <el-table-column label="被考核人" prop="user.nickName" width="460"></el-table-column> 
        <el-table-column label="考核表权重分值" prop="score" width="120"></el-table-column>
        
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
    getPRItemCount
} from "@/api/pas/performanceReviewItem";  //  此处请自行替换地址
import {    
    updatePRStatusById,
    getPRBystatus
} from "@/api/pas/performanceReview";  //  此处请自行替换地址

import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";
import { mapGetters } from "vuex";
import { getDict } from "@/utils/dictionary";
export default {
  name: "confirm",
  mixins: [infoList],
  data() {
    return {
      listApi: getPRItemListByUser,
      type: "",
      multipleSelection: [],
      countData:9,
      dictList:[],
      prData:{
        score:0,
        name:"",
        evaluation:{
          name:"",
          category:"",
          description:"",
        },
        user:{
          nickName:"",
        },
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
      filterDict(status){
        console.log(status)
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
    async confirmKpi(row) {
      const res = await updatePRItemStatusById({
          ID:row.ID,
          status:2,
      })
      if(res.code == 0){
          this.getPRItemListByUser()
          this.$message({
          type: "success",
          message: "确认成功"})
        const count = await getPRItemCount({
            PRId:row.PRId,
            status: 1,
        })
        this.countData = count.data.total;
        if(this.countData == 0){
            updatePRStatusById({
                ID:row.PRId,
                status: 2,
            })
        }
      }
    },
    async rejectKpi(row){
      const res = await updatePRItemStatusById({
          ID:row.ID,
          status:99,
      })
      if(res.code == 0){
          this.getPRItemListByUser()
          this.$message({
          type: "success",
          message: "驳回成功"})
      }
    },
    async getPRBystatus(){
      const res = await getPRBystatus({
          status:3,
          })
      this.prData = res.data.list
    },
  },
  async created() {
    this.getPRBystatus()
    const res = await getDict("PR");
    res.map(item=>item.value = String(item.value))
    this.dictList = res
}
};
</script>

<style>
</style>