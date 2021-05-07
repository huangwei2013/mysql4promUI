$(function () {

    function addButtonFunc(value, row, index) {
        return [
                '<div id="editButton"   class="badge badge-success m-1"><i class="far fa-edit mr-1"></i>编辑</div>',
                '<div id="deleteButton"  class="badge badge-danger m-1"> <i class="fas fa-times mr-1"></i>删除</div>',
            ].join('');
    }

    window.operateEvents = {
        'click #editButton': function (e, value, row, index) {
            vm.update(row.Id);
        }, 'click #deleteButton': function (e, value, row, index) {
            vm.del(row.Id);
        }
    };

    $('#table').bootstrapTable({
        url: '/rule/labels/page',
        method: "GET",
        striped: true,
        cache: false,
        pagination: true,
        pageList: [20, 40, 60, 100],
        pageSize: 20,
        pageNumber: 1,
        sortName: "Id",
        sortOrder: "desc",
        sidePagination: 'server',
        search: false,
        uniqueId: "Id",
        silent: true,
        classes: "table table-hover",
        paginationHAlign: "left",
        paginationDetailHAlign: "right",
        queryParams: queryParams,
        responseHandler: function (res) {
        return {
            "total": res.data.form.TotalSize,
            "rows": res.data.list
        };

        },
        onLoadSuccess: function () {

        },
        onLoadError: function () {
            alert("数据加载失败！");
        },
        columns: [{
                checkbox: true,
                visible: true
                },
                 {
                    field: 'Id',
                    title: 'ID'
                },
                 {
                    field: 'RuleId',
                    title: '规则ID'
                },
                 {
                    field: 'LabelKey',
                    title: '规则标签名'
                },
                 {
                    field: 'LabelValue',
                    title: '规则标签值'
                },
                 {
                    field: 'State',
                    title: '状态'
                },
                 {
                    field: 'CreatedAt',
                    title: '创建时间'
                },
                 {
                    field: 'UpdatedAt',
                    title: '更新时间'
                },
                {
                    field: 'operate',
                    title: '操作',
                    events: operateEvents,
                    formatter: addButtonFunc 
                }
            ]
    });

    function queryParams(params) {
        var temp = {
            offset: params.offset,
            limit: params.limit,
            search: $(".search-input").val(),
            rows: params.limit,
            page: (params.offset / params.limit) + 1,
            sort: params.sort,
            sortOrder: params.order
        };
        return temp;
    };
    $("#search-btn").click(function () {
        $('#table').bootstrapTable('refresh');
    });

});

var vm = new Vue({
    el: '#rrapp',
    data: {
        showList: true,
        title: null,
        trulelabels:{
        }
    },
    methods: {
        query: function () {
            vm.reload();
        },
        add: function () {
            vm.showList = false;
            vm.title = "新增";
            vm.trulelabels= {};
         },
        update: function (Id) {
            if (Id == null) {
                return;
            }
            vm.showList = false;
            vm.title = "修改";
            vm.getInfo(Id)
        },
    saveOrUpdate: function (event) {
        $('#btnSaveOrUpdate').button('loading').delay(1000).queue(function () {
            var url = vm.trulelabels.Id ==null ? "/rule/labels/save" : "/rule/labels/update";
            $.ajax({
                type: "POST",
                url: url,
                data: vm.trulelabels,
                success: function (r) {
                    if (r.code === 0) {
                        swal({
                            text: "操作成功",
                            icon: "success",
                            buttons: false,
                            timer: 2000,
                        });
                        vm.reload();
                        $('#btnSaveOrUpdate').button('reset');
                        $('#btnSaveOrUpdate').dequeue();
                    } else {
                        swal({
                            text: r.msg,
                            icon: "error",
                            buttons: false,
                            timer: 2000,
                        });
                        $('#btnSaveOrUpdate').button('reset');
                        $('#btnSaveOrUpdate').dequeue();
                    }
                }
        })
            ;
        });
    },
    del: function (Id) {
        var Ids = [Id];
        var lock = false;
        top.swal({
        title: "确定要删除该记录?",
        icon: "warning",
        buttons: ["取消", "确定"],
        closeModal: true,
    }).then((isConfirm) => {
            if (isConfirm) {
            top.swal.close();
            if (!lock) {
                lock = true;
                $.ajax({
                    type: "POST",
                    url:  + "/rule/labels/delete",
                    data: {ids: Ids},
                    success: function (r) {
                        if (r.code == 0) {
                            swal({
                                text: "删除成功",
                                icon: "success",
                                buttons: false,
                                timer: 2000,
                            });
                            $('#table').bootstrapTable('refresh');
                        } else {
                            swal({
                                text: r.msg,
                                icon: "error",
                                buttons: false,
                                timer: 2000,
                            });
                        }
                    }
                });
            }
        }else {
            top.swal.close();
        }
    });
    },
    bulkdel: function (event) {
        var Ids = getSelectedRows();
        if (Ids == null || Ids.length == 0) {
            return;
        }
        var lock = false;
        top.swal({
            title: "确定要删除该记录?",
            icon: "warning",
            buttons: ["取消", "确定"],
            closeModal: true,
        }).then((isConfirm) => {
            if (isConfirm) {
                top.swal.close();
                if (!lock) {
                    lock = true;
                    $.ajax({
                        type: "POST",
                        url: "/rule/labels/delete",
                        data: {ids: Ids},
                        success: function (r) {
                            if (r.code == 0) {
                                swal({
                                    text: "删除成功",
                                    icon: "success",
                                    buttons: false,
                                    timer: 2000,
                                });
                                $('#table').bootstrapTable('refresh');
                            } else {
                                swal({
                                    text: r.msg,
                                    icon: "error",
                                    buttons: false,
                                    timer: 2000,
                                });
                            }
                        }
                    });
                }
            }else {
                top.swal.close();
            }
        });
    },
    getInfo: function (Id) {
        $.get("/rule/labels/get/" +Id, function (r) {
            vm.trulelabels= r.data;
        });
    },
    reload: function (event) {
        vm.showList = true;
        $('#table').bootstrapTable('refresh');
    }
}
})
;