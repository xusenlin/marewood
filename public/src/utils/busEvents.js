//小项目还达不到使用redux。在这里维护部分全局的状态
const BusEvents = {
  search: {
    setSearchKeyword: keyword => keyword, //设置全局的关键字
    getSearchKeyword: () => {} //获取全局的关键字
  },
  updateCurrentTable: () => {}, //全局更新当前显示的表格到第一页
  pushNotice: () => {} //push一条消息到数组前面
};

// console.log(BusEvents);
export default BusEvents;
