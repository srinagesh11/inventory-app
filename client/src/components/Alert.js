export default function Alert(props) {
    const { message, type } = props;

    return (
        <div>
            {
                message  &&
                <div className={`alert alert-${type} alert-dismissible fade show`} role="alert">
                    {message}
                    <button type="button" className="btn-close" data-bs-dismiss="alert" aria-label="Close"
                        onClick={props.onClose}></button>
                </div>
            }
        </div>
    )
};

