<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">          
        <!-- <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item> -->
        <el-form-item>
          <el-button @click="openDialog" type="primary">新增方案</el-button>
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
      row-key="ID"
      ref="multipleTable"
      stripe
      style="width: 100%"
      tooltip-effect="dark"
    >
    <el-table-column type="selection" width="55"></el-table-column>
    <!-- <el-table-column label="日期" width="180">
         <template slot-scope="scope">{{scope.row.CreatedAt|formatDate}}</template>
    </el-table-column> -->
    <el-table-column label="方案ID" prop="ID" width="120"></el-table-column> 
    <el-table-column label="方案名称" prop="Name" width="120"></el-table-column> 
    
    <!-- <el-table-column label="方案类型" prop="Category" width="120"></el-table-column>  -->
    
    <el-table-column label="方案状态" prop="Status" width="120"></el-table-column> 
    
    <el-table-column label="方案描述" prop="Description" width="120"></el-table-column> 
    
    <el-table-column label="方案总分" prop="Score" width="120"></el-table-column> 
    
      <el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-button class="table-button" @click="updateEvaluation(scope.row)" size="small" type="primary" icon="el-icon-edit">变更</el-button>
          <el-button class="table-button" @click="opdendrawer(scope.row)" size="small" type="primary" icon="el-icon-edit">添加指标</el-button>
          <el-popover placement="top" width="160" v-model="scope.row.visible">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin: 0">
              <el-button size="mini" type="text" @click="scope.row.visible = false">取消</el-button>
              <el-button type="primary" size="mini" @click="deleteEvaluation(scope.row)">确定</el-button>
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
        
         <el-form-item label="方案名称:">
            <el-input v-model="formData.Name" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="方案类型:">
            <el-input v-model="formData.Category" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <!-- <el-form-item label="方案状态:">
            <el-input v-model="formData.Status" clearable placeholder="请输入" ></el-input>
      </el-form-item> -->
       
         <el-form-item label="方案描述:">
            <el-input v-model="formData.Description" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="方案总分:">
            <el-input v-model="formData.Score" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       </el-form>
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button @click="enterDialog" type="primary">确 定</el-button>
      </div>
    </el-dialog>

    <el-drawer :visible.sync="drawer" :with-header="false" size="60%" title="方案配置" v-if="drawer">
      <el-tabs :before-leave="autoEnter" class="role-box" type="border-card">
        <el-tab-pane label="添加指标">
          <Kpis :evaluation="tableData" :row="activeRow" ref="kpis" />
        </el-tab-pane>
      </el-tabs>
    </el-drawer>

  </div>
</template>

<script>
import {
    createEvaluation,
    deleteEvaluation,
    deleteEvaluationByIds,
    updateEvaluation,
    findEvaluation,
    getEvaluationList
} from "@/api/pas/evaluation";  //  此处请自行替换地址
import Kpis from "@/view/pas/evaluation/components/kpis";
import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";
export default {
  name: "Evaluation",
  mixins: [infoList],
  data() {
    return {
      drawer: false,
      activeRow: {},
      listApi: getEvaluationList,
      dialogFormVisible: false,
      kpiFormVisible: false,
      visible: false,
      type: "",
      deleteVisible: false,
      multipleSelection: [],formData: {
            ID:"",
            Name:"",
            Category:"",
            Status:"",
            Description:"",
            Score: "",
            Kpis:0,
            
      },
    };
  },
  components: {
    Kpis
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
    autoEnter(activeName, oldActiveName) {
      const paneArr = ["kpis"];
      if (oldActiveName) {
        if (this.$refs[paneArr[oldActiveName]].needConfirm) {
          this.$refs[paneArr[oldActiveName]].enterAndNext();
          this.$refs[paneArr[oldActiveName]].needConfirm = false;
        }
      }
    },
    opdendrawer(row) {
      this.drawer = true;
      this.activeRow = row;
    },
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
        const res = await deleteEvaluationByIds({ ids })
        if (res.code == 0) {
          this.$message({
            type: 'success',
            message: '删除成功'
          })
          this.deleteVisible = false
          this.getTableData()
        }
      },
    async updateEvaluation(row) {
      const res = await findEvaluation({ ID: row.ID });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.reEvaluation;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
          Name:"",
          Category:"",
          Status:"",
          Description:"",
          Score:"",
          
      };
    },
    async deleteEvaluation(row) {
      this.visible = false;
      const res = await deleteEvaluation({ ID: row.ID });
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
          res = await createEvaluation({...this.formData,Score:Number(this.formData.Score)});
          break;
        case "update":
          res = await updateEvaluation({...this.formData,Score:Number(this.formData.Score)});
          break;
        default:
          res = await createEvaluation({...this.formData,Score:Number(this.formData.Score)});
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

<style lang="scss">

.role-box {
  .el-tabs__content {
    height: calc(100vh - 150px);
    overflow: auto;
  }
}
</style>