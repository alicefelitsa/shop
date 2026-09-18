<template>
  <div class="content">
    <el-card shadow="always">
      <div slot="header">

        <!--搜索栏-->
        <div class="queryForm">
          <el-form :inline="true" :model="where" class="query-form-inline" size="small">
            <el-form-item label="姓名">
              <el-input v-model="where.name" placeholder="请输入" clearable class="queryElInput"></el-input>
            </el-form-item>
            <el-form-item label="邮箱 / WhatsApp">
              <el-input v-model="where.email" placeholder="请输入" clearable class="queryElInput"></el-input>
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
        <el-button type="danger" size="small" icon="el-icon-delete" @click="del">删除</el-button>
        <el-button size="small" icon="el-icon-download" @click="handleExport">导出</el-button>
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
        <el-table-column prop="name" label="姓名" min-width="100px">
          <template v-slot="{row}">
            {{ row.name }}
          </template>
        </el-table-column>
        <el-table-column prop="email" label="邮箱 / WhatsApp" min-width="160px">
          <template v-slot="{row}">
            {{ row.email }}
          </template>
        </el-table-column>
        <el-table-column prop="total_qty" label="总件数" width="80px" align="center">
          <template v-slot="{row}">
            {{ row.total_qty }}
          </template>
        </el-table-column>
        <el-table-column prop="remark" label="备注" min-width="160px" show-overflow-tooltip>
          <template v-slot="{row}">
            {{ row.remark }}
          </template>
        </el-table-column>
        <el-table-column prop="ip" label="IP" min-width="110px">
          <template v-slot="{row}">
            {{ row.ip }}
          </template>
        </el-table-column>
        <el-table-column prop="ip_address" label="IP地址" min-width="140px" show-overflow-tooltip>
          <template v-slot="{row}">
            {{ row.ip_address }}
          </template>
        </el-table-column>
        <el-table-column prop="ctime" label="时间" min-width="130px">
          <template v-slot="{row}">
            {{ row.ctime }}
          </template>
        </el-table-column>
        <el-table-column label="操作" align="center" width="100px">
          <template v-slot="{row}">
            <el-button size="mini" @click="view(row)">查看</el-button>
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

    <!--查看购物意向详情-->
    <el-dialog title="购物意向详情" :close-on-click-modal="false" :visible.sync="dialogVisible" width="900px"
               top="6vh" custom-class="intent-detail-dialog">
      <el-descriptions v-if="viewData.id" :column="2" border>
        <el-descriptions-item label="姓名:">{{ viewData.name }}</el-descriptions-item>
        <el-descriptions-item label="邮箱 / WhatsApp:">{{ viewData.email }}</el-descriptions-item>
        <el-descriptions-item label="总件数:">{{ viewData.total_qty }}</el-descriptions-item>
        <el-descriptions-item label="时间:">{{ viewData.ctime }}</el-descriptions-item>
        <el-descriptions-item label="IP:">{{ viewData.ip }}</el-descriptions-item>
        <el-descriptions-item label="IP地址:">{{ viewData.ip_address }}</el-descriptions-item>
        <el-descriptions-item label="备注:" :span="2">
          <div style="white-space: pre-wrap;">{{ viewData.remark || '—' }}</div>
        </el-descriptions-item>
      </el-descriptions>

      <!--意向商品明细-->
      <h4 class="intent-items-title">意向商品</h4>
      <el-table :data="viewItems" :border="true" size="small">
        <el-table-column prop="name" label="商品名称" min-width="140px"></el-table-column>
        <el-table-column prop="itemNo" label="货号" width="90px">
          <template v-slot="{row}">
            {{ row.itemNo || '—' }}
          </template>
        </el-table-column>
        <el-table-column prop="spec" label="规格" min-width="140px">
          <template v-slot="{row}">
            {{ row.spec || '—' }}
          </template>
        </el-table-column>
        <el-table-column prop="price" label="单价" min-width="110px">
          <template v-slot="{row}">
            {{ row.price || '—' }}
          </template>
        </el-table-column>
        <el-table-column prop="qty" label="数量" width="70px" align="center"></el-table-column>
      </el-table>

      <div slot="footer">
        <el-button size="small" @click="dialogVisible=false">关闭</el-button>
      </div>
    </el-dialog>

  </div>
</template>

<script>
import {DelCartIntent, GetCartIntent} from "@/api/intent";
import {exportToExcel} from "@/utils/excel";

export default {
  name: "intent",
  data() {
    return {
      tableData: [],
      totalData: 0,
      pageSizes: [30, 50, 100, 200],
      multipleSelection: [],
      where: {
        name: '',
        email: '',
        page: 1,
        limit: 30,
      },
      visitorLoading: false,
      dialogVisible: false,
      viewData: {},
      viewItems: []
    }
  },
  mounted() {
    this.where.limit = this.pageSizes[0]
    this.getIntent()
  },
  methods: {
    //获取客户购物意向
    async getIntent() {
      this.visitorLoading = true;
      setTimeout(async () => {
        try {
          let data = await GetCartIntent({...this.where})
          this.tableData = data.data;
          this.totalData = data.count
        } catch (e) {
          this.$message.error(e.message);
        } finally {
          this.visitorLoading = false;
        }
      }, 200)
    },
    //解析意向商品 JSON
    parseItems(raw) {
      try {
        const arr = JSON.parse(raw || '[]')
        return Array.isArray(arr) ? arr : []
      } catch (e) {
        return []
      }
    },
    //商品明细摘要（用于列表导出）
    itemsSummary(raw) {
      return this.parseItems(raw).map(it => {
        const parts = [it.name || '']
        if (it.itemNo) parts.push(it.itemNo)
        if (it.spec) parts.push(it.spec)
        if (it.price) parts.push(it.price)
        return parts.join('/') + ' x' + (it.qty || 0)
      }).join('; ')
    },
    //查询
    search() {
      this.where.page = 1
      this.getIntent()
    },
    //重置搜索条件
    reset() {
      this.where.name = ''
      this.where.email = ''
      this.where.page = 1
      this.getIntent()
    },
    //页数
    handleSizeChange(val) {
      this.where.limit = val
      this.getIntent()
    },
    //页码
    handleCurrentChange(val) {
      this.where.page = val
      this.getIntent()
    },
    //多选
    handleSelectionChange(val) {
      this.multipleSelection = []
      val.forEach(item => {
        this.multipleSelection.push(item.id);
      })
    },
    //查看意向详情
    view(row) {
      this.viewData = {...row}
      this.viewItems = this.parseItems(row.items)
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
          let message = await DelCartIntent(ids)
          this.$message.success(message)
          await this.getIntent()
        } catch (e) {
          this.$message.error(e.message);
        }
      }).catch(_ => {
      });
    },
    //导出表格
    async handleExport() {
      const headers = ['姓名', '邮箱 / WhatsApp', '总件数', '意向商品', '备注', 'IP', 'IP地址', '时间']
      const rows = this.tableData.map(item => [
        item.name || '',
        item.email || '',
        item.total_qty || 0,
        this.itemsSummary(item.items),
        item.remark || '',
        item.ip || '',
        item.ip_address || '',
        item.ctime || ''
      ])
      await exportToExcel(headers, rows, '客户购物意向表')
    }
  }
}
</script>

<style scoped>
/* 意向详情弹窗：表格内容统一 14px 微软雅黑 */
::v-deep .intent-detail-dialog .el-descriptions-item__label,
::v-deep .intent-detail-dialog .el-descriptions-item__content {
  font-size: 14px;
  font-family: "Microsoft YaHei", "微软雅黑", sans-serif;
}

/* 标签列不换行，宽度随内容自适应，文字右对齐 */
::v-deep .intent-detail-dialog .el-descriptions-item__label {
  white-space: nowrap;
  width: 1px;
  text-align: right;
}

/* 弹窗内容区域内边距 */
::v-deep .intent-detail-dialog .el-dialog__body {
  padding: 20px 20px !important;
}

/* 意向商品小标题 */
.intent-items-title {
  margin: 18px 0 10px;
  font-size: 14px;
  font-weight: normal;
  color: #303133;
}
</style>
