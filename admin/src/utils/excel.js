import ExcelJS from 'exceljs'
import {saveAs} from 'file-saver'

/**
 * 导出 Excel（表头粗体 + 冻结首行 + 左对齐 + 自动列宽）
 */
export async function exportToExcel(headers, rows, fileName = '导出数据') {
    const workbook = new ExcelJS.Workbook()
    const worksheet = workbook.addWorksheet('Sheet1')

    // 添加表头
    worksheet.addRow(headers)

    // 添加数据行
    rows.forEach(row => {
        worksheet.addRow(row)
    })

    // 设置列宽
    worksheet.columns = autoColWidth([headers, ...rows])

    // 设置表头样式（第一行）
    const headerRow = worksheet.getRow(1)
    headerRow.font = {
        name: 'Microsoft YaHei',
        size: 11,
        bold: true
    }
    headerRow.alignment = {horizontal: 'left', vertical: 'middle'}

    // 设置数据行样式（从第2行开始）
    for (let i = 2; i <= worksheet.rowCount; i++) {
        const row = worksheet.getRow(i)
        row.eachCell({includeEmpty: true}, (cell) => {
            cell.font = {
                name: 'Microsoft YaHei',
                size: 11
            }
            cell.alignment = {horizontal: 'left', vertical: 'middle'}
        })
    }

    // 冻结首行
    worksheet.views = [
        {state: 'frozen', ySplit: 1}
    ]

    // 导出文件
    const buffer = await workbook.xlsx.writeBuffer()
    saveAs(new Blob([buffer]), `${fileName}.xlsx`)
}

/**
 * 计算字符串显示宽度（中文字符算2，英文算1）
 */
function calcStrWidth(str) {
    let w = 0
    for (const ch of String(str)) {
        w += ch.charCodeAt(0) > 255 ? 2 : 1
    }
    return w
}

/**
 * 自动计算列宽
 */
function autoColWidth(data) {
    if (!data || data.length === 0) return []
    const colCount = data[0].length
    const widths = new Array(colCount).fill(0)

    for (let row of data) {
        for (let c = 0; c < colCount; c++) {
            const val = row[c] == null ? '' : String(row[c])
            const w = calcStrWidth(val)
            if (w > widths[c]) widths[c] = w
        }
    }
    // 字体需要更大的边距，从+4改为+6，留更多空间
    return widths.map(w => ({width: Math.min((w + 6), 80)}))
}


