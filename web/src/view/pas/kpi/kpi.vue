<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">        
        <!-- <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item> -->
        <el-form-item>
          <el-button @click="openDialog" type="primary">新增指标表</el-button>
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
    
    <el-table-column label="指标名称" prop="Name" width="120"></el-table-column> 
    
    <el-table-column label="指标说明" prop="Description" width="360" type="textarea"></el-table-column> 
    
    <el-table-column label="指标状态" prop="Status" width="120"></el-table-column> 
    
    <el-table-column label="指标算法" prop="Category" width="360" type="textarea"></el-table-column> 

    <el-table-column label="标签名称">
      <template slot-scope="scope">
        <span v-for="(item,index) in scope.row.Tags"
        :key="index">{{item.Name}}</span>
      </template>
    </el-table-column>
    <el-table-column label="标签类型">
      <template slot-scope="scope">
        <span v-for="(item,index) in scope.row.Tags"
        :key="index">{{item.Category}}</span>
      </template>
    </el-table-column>
      <el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-button class="table-button" @click="updateKpi(scope.row)" size="small" type="primary" icon="el-icon-edit">编辑指标</el-button>
          <el-popover placement="top" width="160" v-model="scope.row.visible">
            <p>确定要删除本条指标吗？</p>
            <div style="text-align: right; margin: 0">
              <el-button size="mini" type="text" @click="scope.row.visible = false">取消</el-button>
              <el-button type="primary" size="mini" @click="deleteKpi(scope.row)">确定</el-button>
            </div>
            <el-button type="danger" icon="el-icon-delete" size="mini" slot="reference">删除指标</el-button>
          </el-popover>
          <el-popover placement="top" width="160" v-model="scope.row.vis">
            <p>确定要清空标签吗？</p>
            <div style="text-align: right; margin: 0">
              <el-button size="mini" type="text" @click="scope.row.visible = false">取消</el-button>
              <el-button type="primary" size="mini" @click="removeKpiTags(scope.row)">确定</el-button>
            </div>
            <el-button type="danger" icon="el-icon-delete" size="mini" slot="reference">清除标签</el-button>
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
         <el-form-item label="指标名称:">
            <el-input v-model="formData.Name" clearable placeholder="请输入"></el-input>
      </el-form-item>
         <el-form-item label="指标说明:">
            <el-input v-model="formData.Description" clearable placeholder="请输入"  type="textarea"
        :autosize="{minRows: 4, maxRows: 4}"></el-input>
      </el-form-item>
         <el-form-item label="指标状态:">
            <el-input v-model="formData.Status" clearable placeholder="请输入"></el-input>
      </el-form-item>      
         <el-form-item label="指标类型:">
            <el-input v-model="formData.Category" clearable placeholder="请输入"  type="textarea"
        :autosize="{minRows: 4, maxRows: 4}"></el-input>
      </el-form-item>
          <el-form-item label="下拉选择" prop="Tags"> 
            <el-select v-model="formData.Tags" placeholder="请选择下拉选择" clearable
              :style="{width: '100%'}">
              <el-option v-for="(item, index) in tagData" :key="index"
                :value="item.ID" :disabled="item.disabled"
                >{{item.Name}}</el-option>
            </el-select>
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
    createKpi,
    deleteKpi,
    deleteKpiByIds,
    updateKpi,
    findKpi,
    getKpiList,
    removeKpiTags
} from "@/api/pas/kpi";  //  此处请自行替换地址
import{
    getTagList,
}from "@/api/pas/tag"; 
import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";
export default {
  name: "kpi",
  mixins: [infoList],
  data() {
    return {
      listApi: getKpiList,
      dialogFormVisible: false,
      visible: false,
      type: "",
      deleteVisible: false,
      multipleSelection: [],formData: {
            Name:"",
            Description:"",
            Status:"",
            Category:"",
            Tags: 1,
      },
      tagData:{
           Name:"",
           ID:"",
           Category:"",
           Parentid:"",
      }, 
      tagColumn:{
           Name:"",
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
        const res = await deleteKpiByIds({ ids })
        if (res.code == 0) {
          this.$message({
            type: 'success',
            message: '删除成功'
          })
          this.deleteVisible = false
          this.getTableData()
        }
      },
    
    async updateKpi(row) {
      const res = await findKpi({ ID: row.ID });
      //const tag = await findTag({ ID: row.ID });
      const tags = await getTagList();
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.reKpi;
        this.tagData = tags.data.list;
        //this.tagColumn = tag.data.reTag;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
          Name:"",
          Description:"",
          Status:"",
          Category:"",
          Tags: 1,
      };
    },
    async deleteKpi(row) {
      this.visible = false;
      const res = await deleteKpi({ ID: row.ID });
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "删除成功"
        });
        this.getTableData();
      }
    },
    async removeKpiTags(row) {
      this.visible = false;
      const res = await removeKpiTags({ ID: row.ID });
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "清除成功"
        });
        this.getTableData();
      }
    },
    async enterDialog() {
      let res;
      switch (this.type) {
        case "create":
          res = await createKpi({...this.formData,Tags:[{ID:this.formData.Tags}]});
          break;
        case "update":
          res = await updateKpi({...this.formData,Tags:[{ID:this.formData.Tags}]});
          break;
        default:
          res = await createKpi({...this.formData,Tags:[{ID:this.formData.Tags}]});
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
    async openDialog() {
      const tags = await getTagList();
      this.tagData = tags.data.list;
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

.el-table .cell {
  white-space: pre-line;
}
</style>