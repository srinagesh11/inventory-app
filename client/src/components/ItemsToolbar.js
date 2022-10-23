import { useEffect, useState } from "react";

export default function ItemsToolbar(props) {
    const [showForm, setShowForm] = useState(false);
    const [id, setId] = useState("");
    const [name, setName] = useState("");
    const [quantity, setQuantity] = useState("");
    const [unitPrice, setUnitPrice] = useState("");

    useEffect(() => {
        if (props.selectedItem) {
            setShowForm(true);
            setId(props.selectedItem.id);
            setName(props.selectedItem.name);
            setQuantity(props.selectedItem.quantity);
            setUnitPrice(props.selectedItem.unit_price);
        }
    }, [props]);

    // function to handle the form submission
    const submitHandler = (event) => {
        event.preventDefault();

        let itemId = parseInt(id);
        let itemQuantity = parseInt(quantity);
        let itemUnitPrice = parseFloat(unitPrice);


        if (itemId <= 0) {
            props.setErr("Please enter an ID for the item");
            return;
        }
        if (name === "") {
            props.setErr("Please enter a name for the item");
            return;
        }
        if (itemQuantity <= 0) {
            props.setErr("Please enter a quantity for the item");
            return;
        }
        if (itemUnitPrice <= 0) {
            props.setErr("Please enter a unit price for the item");
            return;
        }

        let func = props.selectedItem ? props.onEditItem : props.onAddItem;
        console.log(func);
        func({
            id: itemId,
            name: name,
            quantity: itemQuantity,
            unit_price: itemUnitPrice
        }, function(ok) {
            if (ok) {
                setId("");
                setName("");
                setQuantity("");
                setUnitPrice("");
                setShowForm(false);
            }
        });
    };


    return (
        <div className="toolbar">
            <div className="row">
                <div className="col pb-1">
                    <h4>Your Inventory</h4>
                </div>
                <div className="col" style={{textAlign: 'right'}}>
                    <button className="btn btn-success" style={{marginRight: '10px'}}
                        onClick={event => setShowForm(prev => !prev)}
                    >{showForm ? "Hide Form" : "Add Item"}</button>
                    <button className="btn btn-secondary"
                        onClick={event => props.downloadCSV()}
                    >Download CSV</button>
                </div>
            </div>
            <hr/>
            {showForm && <div className="border rounded p-3 mb-3 shadow p-3 mb-5 bg-body">
                <p>Fill out this form to add a new item to your inventory</p>
                <hr/>
                <form onSubmit={submitHandler}>
                    <div className="row mb-3">
                        <div className="col">
                            <label htmlFor="itemId" className="form-label">Id</label>
                            <input type="number" step="1" min="1" className="form-control" id="itemId"
                                value={id}
                                onChange={event => setId(event.target.value)} />
                        </div>
                        <div className="col">
                            <label htmlFor="itemName" className="form-label">Item Name</label>
                            <input type="text" className="form-control" id="itemName" required
                                value={name}
                                onChange={event => setName(event.target.value)}/>
                        </div>
                    </div>
                    <div className="row mb-3">
                        <div className="col">
                            <label htmlFor="itemQuantity" className="form-label">Quantity</label>
                            <input type="number" step="1" min="1" className="form-control" id="itemQuantity"
                                value={quantity}
                                onChange={event => setQuantity(event.target.value)}/>
                        </div>
                        <div className="col">
                            <label htmlFor="itemUnitPrice" className="form-label">Unit Price</label>
                            <input type="number" className="form-control" id="itemUnitPrice" step={0.01}
                                value={unitPrice}
                                onChange={event => setUnitPrice(event.target.value)}/>
                        </div>
                    </div>
                    <button type="submit" className="btn btn-primary w-100">Submit</button>
                </form>                
            </div>}
        </div>
    )
}