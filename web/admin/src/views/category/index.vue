<template>
  <div class="content">
    <el-card shadow="always">
      <div slot="header">

        <!--搜索栏-->
        <div class="queryForm">
          <el-form :inline="true" :model="where" class="query-form-inline" size="small">
            <el-form-item label="名称">
              <el-input v-model="where.name" placeholder="请输入" clearable class="queryElInput"></el-input>
            </el-form-item>
            <el-form-item>
              <el-button type="primary" @click="search">查询</el-button>
              <el-button icon="el-icon-refresh" @click="reset">重置</el-button>
            </el-form-item>
          </el-form>
        </div>
      </div>

      <!--工具栏-->
      <div class="toolbar">
        <el-button type="primary" size="small" icon="el-icon-plus" @click="add">添加</el-button>
        <el-button type="danger" size="small" icon="el-icon-delete" @click="del">删除</el-button>
      </div>

      <!--数据表格-->
      <el-table class="tableData" :data="tableData" :highlight-selection-row="true" height="calc(100vh - 182px)"
                :border="true"
                v-loading="visitorLoading" @selection-change="handleSelectionChange">
        <el-table-column type="selection" width="55" align="center"></el-table-column>
        <el-table-column prop="id" label="ID" width="80px">
          <template v-slot="{row}">
            {{ row.id }}
          </template>
        </el-table-column>
        <el-table-column prop="name" label="分类名称" min-width="160px" show-overflow-tooltip>
          <template v-slot="{row}">
            {{ row.name }}
          </template>
        </el-table-column>
        <el-table-column label="操作" align="center" width="100px">
          <template v-slot="{row}">
            <el-button size="mini" @click="edit(row)">编辑</el-button>
          </template>
        </el-table-column>
      </el-table>

      <!--表格分页-->
      <div style="margin-top: 10px; text-align: center;" class="currentPage">
        <el-pagination
            @size-change="handleSizeChange"
            @current-change="handleCurrentChange"
            :current-page="where.page"
            :page-sizes="pageSizes"
            :page-size="where.limit"
            background
            layout="total, sizes, prev, pager, next, jumper"
            :total="totalData">
        </el-pagination>
      </div>

    </el-card>

    <!--添加/编辑分类-->
    <el-dialog :title="form.id?'修改分类':'添加分类'" :visible.sync="dialogVisible" width="400px"
               :close-on-click-modal="false">
      <el-form ref="form" :model="form" label-width="80px">
        <el-form-item label="名称">
          <el-input v-model="form.name" placeholder="请输入分类名称"></el-input>
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button @click="dialogVisible=false" size="small">取 消</el-button>
        <el-button type="primary" @click="save" size="small">确 定</el-button>
      </span>
    </el-dialog>

  </div>
</template>

<script>
import {AddCategory, DelCategory, GetCategory, SaveCategory} from "@/api/category";

export default {
  name: "category",
  data() {
    return {
      tableData: [],
      totalData: 0,
      pageSizes: [30, 50, 100, 200],
      multipleSelection: [],
      where: {
        name: '',
        page: 1,
        limit: 30,
      },
      visitorLoading: false,
      dialogVisible: false,
      form: {
        id: '',
        name: ''
      }
    }
  },
  mounted() {
    this.where.limit = this.pageSizes[0]
    this.getCategory()
  },
  methods: {
    //获取分类列表
    async getCategory() {
      this.visitorLoading = true;
      setTimeout(async () => {
        try {
          let data = await GetCategory({...this.where})
          this.tableData = data.data;
          this.totalData = data.count
        } catch (e) {
          this.$message.error(e.message);
        } finally {
          this.visitorLoading = false;
        }
      }, 200)
    },
    //查询
    search() {
      this.where.page = 1
      this.getCategory()
    },
    //重置搜索条件
    reset() {
      this.where.name = ''
      this.where.page = 1
      this.getCategory()
    },
    //页数
    handleSizeChange(val) {
      this.where.limit = val
      this.getCategory()
    },
    //页码
    handleCurrentChange(val) {
      this.where.page = val
      this.getCategory()
    },
    //多选
    handleSelectionChange(val) {
      this.multipleSelection = []
      val.forEach(item => {
        this.multipleSelection.push(item.id);
      })
    },
    //添加
    add() {
      this.form = {id: '', name: ''}
      this.dialogVisible = true
    },
    //编辑
    edit(row) {
      this.form = {...row}
      this.dialogVisible = true
    },
    //添加或修改
    async save() {
      if (!this.form.name.trim()) {
        this.$message.warning("请输入分类名称");
        return
      }
      try {
        const addOrSave = this.form.id ? SaveCategory : AddCategory;
        let message = await addOrSave({...this.form})
        this.$message.success(message)
        this.dialogVisible = false
        await this.getCategory()
      } catch (e) {
        this.$message.error(e.message);
      }
    },
    //删除
    async del() {
      if (this.multipleSelection.length === 0) {
        this.$message.warning("请选择要删除的数据");
        return
      }
      this.$confirm('即将删除，是否继续?').then(async _ => {
        let ids = this.multipleSelection.join(',')
        try {
          let message = await DelCategory(ids)
          this.$message.success(message)
          await this.getCategory()
        } catch (e) {
          this.$message.error(e.message);
        }
      }).catch(_ => {
      });
    }
  }
}
</script>
