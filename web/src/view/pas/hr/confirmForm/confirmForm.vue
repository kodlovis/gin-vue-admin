<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item>
          <el-popover placement="top" v-model="confirmVisible" width="160">
            <p>要批量确认吗？</p>
              <div style="text-align: right; margin: 0">
                <el-button @click="confirmVisible = false" size="mini" type="text">取消</el-button>
                <el-button @click="onConfirm" size="mini" type="primary">确定</el-button>
              </div>
            <el-button size="mini" slot="reference" type="primary">批量确认</el-button>
          </el-popover>
        </el-form-item>
      </el-form>
    </div>
      <el-table
          :data="prData"
          @selection-change="handleSelectionChange"
          border
          ref="multipleTable"
          stripe
          style="width: 100%"
          tooltip-effect="dark"
        >
        <el-table-column type="selection" width="55"></el-table-column>
        
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
    updatePRItemStatusByPrId,
    updatePRItemStatysByIds
} from "@/api/pas/performanceReviewItem";  //  此处请自行替换地址
import {    
    getPRBystatus,
    updatePRStatysByIds,
    updatePRStatusById,
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
      confirmVisible:false,
      page: 1,
      total: 10,
      pageSize: 10,
      dictList:[],formData: [],
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
      const res = await updatePRStatusById({
          ID:row.ID,
          status:5,
      })
      if(res.code == 0){
          this.getPRBystatus()
          this.$message({
          type: "success",
          message: "确认成功"})
          updatePRItemStatusByPrId({PRId:row.ID,status:4})
      }
    },
    async getPRBystatus(page = this.page, pageSize = this.pageSize){
      const res = await getPRBystatus({
          status:3,
          page: page, 
          pageSize: pageSize
          })
      this.total = res.data.total
      this.page = res.data.page
      this.pageSize = res.data.pageSize
      this.prData = res.data.list
    },
    handleSizeChange(val) {
        this.pageSize = val
        this.getPRBystatus()
    },
    handleCurrentChange(val) {
        this.page = val
        this.getPRBystatus()
    },
    async onConfirm(){
      const ids = []
        if(this.multipleSelection.length == 0){
          this.$message({
            type: 'warning',
            message: '请选择要开始确认的考核'
          })
          return
        }
        this.multipleSelection &&
          this.multipleSelection.map(item => {
            ids.push(item.ID)
          })
        const res = await updatePRStatysByIds({ids:ids,status:5})
        if (res.code == 0) {
          this.$message({
            type: 'success',
            message: '发布成功'
          })
          updatePRItemStatysByIds({ids:ids,status:4})
          this.confirmVisible = false
          this.getPRBystatus()
        }
    },
  },
  async created() {
    this.getPRBystatus()
    const res = await getDict("PR");
    res.map(item=>item.value)
    this.dictList = res
}
};
</script>

<style>
</style>