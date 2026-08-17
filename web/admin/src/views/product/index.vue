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
            <el-form-item label="分类">
              <el-select v-model="where.category" placeholder="请选择" clearable class="queryElInput">
                <el-option v-for="item in categoryList" :key="item.id" :label="item.name"
                           :value="item.name"></el-option>
              </el-select>
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
        <el-button type="primary" size="small" icon="el-icon-plus" @click="dialogVisible=true">添加</el-button>
        <el-button type="danger" size="small" icon="el-icon-delete" @click="del">删除</el-button>
      </div>

      <!--数据表格-->
      <el-table class="tableData" :data="tableData" :highlight-selection-row="true" height="calc(100vh - 182px)"
                :border="true"
                v-loading="visitorLoading" @selection-change="handleSelectionChange">
        <el-table-column type="selection" width="55" align="center"></el-table-column>
        <el-table-column prop="id" label="ID" width="60px">
          <template v-slot="{row}">
            {{ row.id }}
          </template>
        </el-table-column>
        <el-table-column prop="album" label="图片" align="center" width="120px">
          <template v-slot="{row}">
            <el-image
                style="height: 80px; width: 80px;"
                :src="row.album"
                fit="contain"
                :preview-src-list="[row.album]"
            >
              <div slot="error" class="image-slot">
                <i class="el-icon-picture-outline"></i>
              </div>
            </el-image>
          </template>
        </el-table-column>
        <el-table-column prop="name" label="名称" min-width="160px" show-overflow-tooltip>
          <template v-slot="{row}">
            {{ row.name }}
          </template>
        </el-table-column>
        <el-table-column prop="category" label="分类" min-width="100px">
          <template v-slot="{row}">
            {{ row.category }}
          </template>
        </el-table-column>
        <el-table-column prop="price" label="价格" min-width="100px">
          <template v-slot="{row}">
            {{ row.price }}
          </template>
        </el-table-column>
        <el-table-column prop="level" label="级别" width="80px">
          <template v-slot="{row}">
            {{ row.level }}
          </template>
        </el-table-column>
        <el-table-column prop="purity" label="纯度" width="80px">
          <template v-slot="{row}">
            {{ row.purity }}
          </template>
        </el-table-column>
        <el-table-column prop="Introduction" label="简介" min-width="180px" show-overflow-tooltip>
          <template v-slot="{row}">
            {{ row.Introduction }}
          </template>
        </el-table-column>
        <el-table-column prop="ctime" label="时间" min-width="130px">
          <template v-slot="{row}">
            {{ row.ctime }}
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

    <!--编辑数据-->
    <save :visible.sync="dialogVisible" :edit-data="editData" :category-list="categoryList"
          @done="getProductList"></save>

  </div>
</template>

<script>
import {DelProduct, GetProductList} from "@/api/productAdmin";
import {GetCategory} from "@/api/category";
import save from "./save";

export default {
  components: {save},
  data() {
    return {
      tableData: [],
      totalData: 0,
      pageSizes: [30, 50, 100, 200],
      multipleSelection: [],
      where: {
        name: '',
        category: '',
        page: 1,
        limit: 30,
      },
      visitorLoading: false,
      dialogVisible: false,
      editData: {},
      categoryList: []
    }
  },
  mounted() {
    this.where.limit = this.pageSizes[0]
    this.getProductList()
    this.getCategoryList()
  },
  methods: {
    //获取产品列表
    async getProductList() {
      this.visitorLoading = true;
      setTimeout(async () => {
        try {
          let data = await GetProductList({...this.where})
          this.tableData = data.data;
          this.totalData = data.count
        } catch (e) {
          this.$message.error(e.message);
        } finally {
          this.visitorLoading = false;
        }
      }, 200)
    },
    //获取分类，供搜索和编辑下拉使用
    async getCategoryList() {
      try {
        let data = await GetCategory({page: 1, limit: 1000})
        this.categoryList = data.data || []
      } catch (e) {
        this.$message.error(e.message);
      }
    },
    //查询
    search() {
      this.where.page = 1
      this.getProductList()
    },
    //重置搜索条件
    reset() {
      this.where.name = ''
      this.where.category = ''
      this.where.page = 1
      this.getProductList()
    },
    //页数
    handleSizeChange(val) {
      this.where.limit = val
      this.getProductList()
    },
    //页码
    handleCurrentChange(val) {
      this.where.page = val
      this.getProductList()
    },
    //多选
    handleSelectionChange(val) {
      this.multipleSelection = []
      val.forEach(item => {
        this.multipleSelection.push(item.id);
      })
    },
    //获取修改数据
    edit(row) {
      this.editData = {...row}
      this.dialogVisible = true;
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
          let message = await DelProduct(ids)
          this.$message.success(message)
          await this.getProductList()
        } catch (e) {
          this.$message.error(e.message);
        }
      }).catch(_ => {
      });
    }
  }
}
</script>

<style>
.image-slot {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100%;
  height: 100%;
}

.image-slot i {
  font-size: 30px;
  color: #909399;
}
</style>
