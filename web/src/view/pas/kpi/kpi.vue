<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-select v-model="searchInfo.status" placeholder="请选择" clearable>
          <el-option
            v-for="item in dictList"
            :key="item.value"
            :label="item.label"
            :value="item.value">
          </el-option>
        </el-select>  
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
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
    
    <el-table-column label="指标名称" prop="name" width="120"></el-table-column> 
    
    <el-table-column label="指标说明" prop="description" width="360" type="textarea"></el-table-column> 
    
    <el-table-column label="指标状态" prop="status" width="120">
      <template slot-scope="scope">
        <!-- <span>{{scope.row.status==0?"评分人确认":""}}</span> -->
        <span>{{filterDict(scope.row.status)}}</span>
      </template>
    </el-table-column>
    
    <el-table-column label="指标算法" prop="category" width="360" type="textarea"></el-table-column> 

    <el-table-column label="标签名称">
      <template slot-scope="scope">
        <span v-for="(item,index) in scope.row.Tags"
        :key="index">{{item.name}}<br/></span>
      </template>
    </el-table-column>
    <el-table-column label="标签类型">
      <template slot-scope="scope">
        <span v-for="(item,index) in scope.row.Tags"
        :key="index">{{item.category}}<br/></span>
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
              <el-button size="mini" type="text" @click="scope.row.vis = false">取消</el-button>
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

    <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="修改指标">
      <el-form :model="formData" label-position="right" label-width="80px" :rules="rules">
         <el-form-item label="指标名称:">
            <el-input v-model="formData.name" clearable placeholder="请输入"></el-input>
      </el-form-item>
         <el-form-item label="指标说明:">
            <el-input v-model="formData.description" clearable placeholder="请输入"  type="textarea"
        :autosize="{minRows: 4, maxRows: 4}"></el-input>
      </el-form-item>
         <el-form-item label="指标状态:">
          <el-select v-model="formData.status" placeholder="请选择">
            <el-option
              v-for="item in dictList"
              :key="item.value"
              :label="item.label"
              :value="item.value">
          </el-option>
        </el-select>  
      </el-form-item>
         <el-form-item label="指标算法:">
            <el-input v-model="formData.category" clearable placeholder="请输入"  type="textarea"
        :autosize="{minRows: 4, maxRows: 4}"></el-input>
      </el-form-item>
          <el-form-item label="添加标签"> 
            <!-- <el-select v-model="formData.Tags" placeholder="请选择需要添加的标签" clearable
              :style="{width: '100%'}">
              <el-option v-for="(item, index) in tagData" 
                :key="index"
                :value="item.ID" 
                :label="item.name"
                :disabled="item.disabled"
                >{{item.name}}</el-option>
            </el-select> -->
            
              <el-cascader
                @change="(val)=>{handleOptionChange(val)}"
                v-model="formData.Tags"
                :options="tagOptions"
                clearable
                :props="{ checkStrictly: true,label:'name',value:'id',multiple: true}"
                filterable
              ></el-cascader>
            <!-- <el-input v-model="formData.Tags"></el-input> -->
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
    findKpi,
    getKpiList,
    removeKpiTags,
    getLastKpi,
    updateKpi
} from "@/api/pas/kpi";  //  此处请自行替换地址
import{
    getTagList,
}from "@/api/pas/tag"; 
import { formatTimeToStr } from "@/utils/date";
import { getDict } from "@/utils/dictionary";
import infoList from "@/mixins/infoList";
import {
    createKpiTag
} from "@/api/pas/kpiTag"; 
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
      statusMap: {},
      dictList:[],
      tagOptions:[],
      multipleSelection: [],formData: {
            name:"",
            description:"",
            status:"",
            category:"",
            Tags: [{id:"",name:""}],
      },
      tagData:{
           name:"",
           ID:"",
           category:"",
           parentid:"",
      }, 
      tagColumn:{
           name:"",
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
      setTagOptions(tagData) {
        this.tagOptions = [];
        this.ids = [];
        this.setTagOptionsData(tagData, this.tagOptions ,this.ids);
      },
      setTagOptionsData(TagData, optionsData ,ids) {
        TagData &&
          TagData.map(item => {
              const option = {
                id: item.ID,
                name: item.name
              };
              optionsData.push(option);
              const idOption = {
                id: item.ID,
              };
              ids.push(idOption)
          });
      },
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
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.reKpi;
        const arr = []
        for (let index = 0; index < this.formData.Tags.length; index++) {
          arr.push(this.formData.Tags[index].ID)
        }
        this.formData.Tags = arr
        //this.tagColumn = tag.data.reTag;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
          name:"",
          description:"",
          status:"",
          category:"",
          Tags: [],
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
      let ref;
      var item = []
      switch (this.type) {
        case "create":
          var re = await createKpi({...this.formData,Tags:[]});
          if(re.code == 0){
            var kpi = await getLastKpi()
            var lastKpi = kpi.data.reKpi
            for (let i = 0; i < this.formData.Tags.length; i++) {
              item.push({
                tagId:Number(this.formData.Tags[i]),
                kpiId:Number(lastKpi.ID),
                })
            } 
            res = await createKpiTag({item})
          }
          break;
        case "update":
          ref = await removeKpiTags({ID: this.formData.ID})
          if (ref.code == 0 && this.formData.Tags != null) {
            for (let i = 0; i < this.formData.Tags.length; i++) {
              item.push({
                tagId:Number(this.formData.Tags[i]),
                kpiId:Number(this.formData.ID),
                })
              }
          createKpiTag({item})
          }
          res = await updateKpi({...this.formData,Tags:[]});
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
    const res = await getDict("kpiLibrary");
    res.map(item=>item.value)
    this.dictList = res
    const tag = await getTagList({ page: 1, pageSize: 999 });
    //载入Users
    this.setTagOptions(tag.data.list);
}
};

</script>

<style>

.el-table .cell {
  white-space: pre-line;
}
</style>