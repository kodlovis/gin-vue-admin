<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">              
        <!-- <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item> -->
        <el-form-item>
          <el-button @click="openDialog" type="primary">人员编辑</el-button>
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
    <!-- <el-table-column label="日期" width="180">
         <template slot-scope="scope">{{scope.row.CreatedAt|formatDate}}</template>
    </el-table-column> -->
    
    <!-- <el-table-column label="EID（用来获得方案表）" prop="EvaluationId" width="120"></el-table-column>  -->
    <el-table-column label="方案名称">
      <template slot-scope="scope">
        <span>{{scope.row.evaluation.Name}}</span>
      </template>
    </el-table-column>
    <el-table-column label="方案类型">
      <template slot-scope="scope">
        <span>{{scope.row.evaluation.Category}}</span>
      </template>
    </el-table-column>
    <el-table-column label="方案描述">
      <template slot-scope="scope">
        <span>{{scope.row.evaluation.Description}}</span>
      </template>
    </el-table-column>
    <!-- <el-table-column label="UID（用来获得User表）" prop="UID" width="120"></el-table-column> 
    
    <el-table-column label="评分人" prop="Scorer" width="120"></el-table-column> 
    
    <el-table-column label="原本的分数" prop="OriginalScore" width="120"></el-table-column>  -->
    
    <el-table-column label="更改后的分数" prop="AdjustedScore" width="120"></el-table-column> 
    
    <el-table-column label="开始日期" prop="StartDate" width="120"></el-table-column> 
    
    <el-table-column label="结束日期" prop="EndingDate" width="120"></el-table-column> 
    
      <el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-button class="table-button" @click="updateAllocation(scope.row)" size="small" type="primary" icon="el-icon-edit">变更</el-button>
          <el-popover placement="top" width="160" v-model="scope.row.visible">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin: 0">
              <el-button size="mini" type="text" @click="scope.row.visible = false">取消</el-button>
              <el-button type="primary" size="mini" @click="deleteAllocation(scope.row)">确定</el-button>
            </div>
            <el-button type="danger" icon="el-icon-delete" size="mini" slot="reference">删除</el-button>
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

    <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="弹窗操作">
      <el-form :model="formData" label-position="right" label-width="80px">
         <el-form-item label="EID（用来获得方案表）:">
            <el-input v-model="formData.EvaluationId" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="UID（用来获得User表）:">
            <el-input v-model="formData.UID" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="评分人:">
            <el-input v-model="formData.Scorer" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="原本分数:">
            <el-input v-model="formData.OriginalScore" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="更改后的分数:">
            <el-input v-model="formData.AdjustedScore" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="开始日期:">
              <el-date-picker type="date" placeholder="选择日期" v-model="formData.StartDate" clearable></el-date-picker>
       </el-form-item>
       
         <el-form-item label="结束日期:">
              <el-date-picker type="date" placeholder="选择日期" v-model="formData.EndingDate" clearable></el-date-picker>
       </el-form-item>
       </el-form>
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button @click="enterDialog" type="primary">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import {
    createAllocation,
    deleteAllocation,
    deleteAllocationByIds,
    updateAllocation,
    findAllocation,
    getAllocationList
} from "@/api/pas/allocation";  //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";
export default {
  name: "Allocation",
  mixins: [infoList],
  data() {
    return {
      listApi: getAllocationList,
      dialogFormVisible: false,
      visible: false,
      type: "",
      deleteVisible: false,
      multipleSelection: [],formData: {
            EvaluationId:"",
            UID:"",
            Scorer:"",
            OriginalScore:"",
            AdjustedScore:"",
            StartDate:new Date(),
            EndingDate:new Date(),
            
      }
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
      //条件搜索前端看此方法
      onSubmit() {
        this.page = 1
        this.pageSize = 10           
        this.getTableData()
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
        const res = await deleteAllocationByIds({ ids })
        if (res.code == 0) {
          this.$message({
            type: 'success',
            message: '删除成功'
          })
          this.deleteVisible = false
          this.getTableData()
        }
      },
    async updateAllocation(row) {
      const res = await findAllocation({ ID: row.ID });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.reAllocation;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
          EvaluationId:"",
          UID:"",
          Scorer:"",
          OriginalScore:"",
          AdjustedScore:"",
          StartDate:new Date(),
          EndingDate:new Date(),
          
      };
    },
    async deleteAllocation(row) {
      this.visible = false;
      const res = await deleteAllocation({ ID: row.ID });
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "删除成功"
        });
        this.getTableData();
      }
    },
    async enterDialog() {
      let res;
      switch (this.type) {
        case "create":
          res = await createAllocation(this.formData);
          break;
        case "update":
          res = await updateAllocation(this.formData);
          break;
        default:
          res = await createAllocation(this.formData);
          break;
      }
      if (res.code == 0) {
        this.$message({
          type:"success",
          message:"创建/更改成功"
        })
        this.closeDialog();
        this.getTableData();
      }
    },
    openDialog() {
      this.type = "create";
      this.dialogFormVisible = true;
    }
  },
  async created() {
    await this.getTableData();
  
}
};
</script>

<style>
</style>