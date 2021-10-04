
let attention = Prompt();
// Example starter JavaScript for disabling form submissions if there are invalid fields
(function () {
    'use strict';
    window.addEventListener('load', function () {
        // Fetch all the forms we want to apply custom Bootstrap validation styles to
        let forms = document.getElementsByClassName('needs-validation');
        // Loop over them and prevent submission
        Array.prototype.filter.call(forms, function (form) {
            form.addEventListener('submit', function (event) {
                if (form.checkValidity() === false) {
                    event.preventDefault();
                    event.stopPropagation();
                }
                form.classList.add('was-validated');
            }, false);
        });
    }, false);
})();
// console.log("Hello from js");
// document.getElementById("colorButton").addEventListener("click", function () {
//     // notieAlert("FUCK YOU!", "success")
//     // sweetAlert("title", "<i>Hello World!</i>", "success", "Okey!")
//     // attention.toast({msg:"Hello World!"});
//     // attention.toast({ msg: "Hello World!", icon: "error" });
//     attention.error({ msg: "hiiiiii!", footer: '<a href="#">This is Link</a>', })
// })


function notieAlert(msg, msgType) {
    notie.alert({
        type: msgType,
        text: msg,
    })
}

function sweetAlert(title, text, icon, confirmButtonText) {
    Swal.fire({
        title: title,
        html: text,
        icon: icon,
        confirmButtonText: confirmButtonText
    })
}


function Prompt() {
    let toast = function (c) {
        const {
            msg = "",
            icon = "success",
            position = "top-end",

        } = c;
        const Toast = Swal.mixin({
            toast: true,
            title: msg,
            position: position,
            icon: icon,
            showConfirmButton: false,
            timer: 3000,
            timerProgressBar: true,
            didOpen: (toast) => {
                toast.addEventListener('mouseenter', Swal.stopTimer)
                toast.addEventListener('mouseleave', Swal.resumeTimer)
            }
        })

        Toast.fire({})
    }

    let success = function (c) {
        const {
            msg = "",
            title = "",
            footer = "",
        } = c;

        Swal.fire({
            icon: 'success',
            title: title,
            text: msg,
            footer: footer,
        })
    }

    let error = function (c) {
        const {
            msg = "",
            title = "",
            footer = "",
        } = c;

        Swal.fire({
            icon: 'error',
            title: title,
            text: msg,
            footer: footer,
        })
    }

    async function custom(c){
        const {
            icon = "",
            msg = "",
            title = "",
            showConfirmButton = true,
        } = c;

        const { value: result} = await Swal.fire({
            icon: icon,
            title: title,
            html: msg,
            backdrop: false,
            focusConfirm: false,
            showCancelButton:true,
            showConfirmButton: showConfirmButton,
            willOpen: () => {
                if (c.willOpen !== undefined){
                    c.willOpen();
                }
            },
            didOpen: () => {
                if (c.didOpen !== undefined){
                    c.didOpen();
                }
            }
        })

        if (result) {
            if (result.dismiss !== Swal.DismissReason.cancel) {
                if (result.value !== "") {
                    if (c.callback !== undefined){
                        c.callback(result);
                    } else {
                        c.callback(false);
                    }
                } else {
                    c.callback(false);
                }
            }
        }
    }

    return {
        toast: toast,
        success: success,
        error: error,
        custom: custom,
    }
}

