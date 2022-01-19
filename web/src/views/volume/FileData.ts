export const columns = [
  {
    title: '名称',
    dataIndex: 'name',
    align: 'left',
    slots: { customRender: 'name' },
  },
  {
    title: '大小',
    dataIndex: 'size',
    width: 150,
    slots: { customRender: 'size' },
  },
  {
    title: '类型',
    dataIndex: 'ext',
    width: 150,
    slots: { customRender: 'ext' },
  },
  {
    title: '修改日期',
    width: 200,
    dataIndex: 'updated_at',
    slots: { customRender: 'updated_at' },
  },
];
