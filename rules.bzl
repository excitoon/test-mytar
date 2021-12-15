def _file_size_impl(ctx):
    out_file = ctx.actions.declare_file("%s.size" % ctx.attr.name)
    ctx.actions.run_shell(
        inputs = [ctx.file.src],
        outputs = [out_file],
        progress_message = "Getting size of %s" % ctx.file.src.short_path,
        command = "wc -c '%s' | awk '{print $1}' | tee '%s'" % (ctx.file.src.path, out_file.path),
    )
    return [DefaultInfo(files = depset([out_file]))]


def _pkg_tar_impl(ctx):
    out_file = ctx.actions.declare_file("%s.tar" % ctx.attr.name)
    ctx.actions.run(
        inputs = ctx.files.srcs,
        outputs = [out_file],
        arguments = [out_file.path, ctx.attr.package_dir] + [src.path for src in ctx.files.srcs],
        progress_message = "Tarring %s" % out_file.short_path,
        executable = ctx.file._mytar
    )
    return [DefaultInfo(files = depset([out_file]))]


file_size = rule(
    implementation = _file_size_impl,
    attrs = {
        "src": attr.label(
            mandatory = True,
            allow_single_file = True,
        ),
    },
)

pkg_tar = rule(
    implementation = _pkg_tar_impl,
    attrs = {
        "srcs": attr.label_list(
            default = [],
            allow_files = True,
        ),
        "_mytar": attr.label(
            allow_single_file = True,
            default = "//:mytar",
        ),
        "package_dir": attr.string(
            default = "",
        ),
    },
)
