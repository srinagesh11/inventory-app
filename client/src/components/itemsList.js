
export default function ItemsList(props) {
    let items = props.items;
    return (
        <div>
            {items.length > 0 && <table className="table">
            <thead className="thead-light">
                <tr>
                    <th scope="col">Id</th>
                    <th scope="col">Name</th>
                    <th scope="col">Quantity</th>
                    <th scope="col">Unit Price</th>
                    <th scope="col" className="text-center">Actions</th>
                </tr>
            </thead>
            <tbody>
                {items.map((item, index) => 
                <tr key={item.id}>
                    <th scope="row">{item.id}</th>
                    <td>{item.name}</td>
                    <td>{item.quantity}</td>
                    <td>{item.unit_price}</td>
                    <td className="text-center">
                        <button className="btn btn-primary btn-sm" style={{marginRight: '10px'}}
                            onClick={event => props.onSelectItemToEdit(item)}>Edit</button>
                        <button className="btn btn-danger btn-sm" 
                            onClick={event => props.onDeleteItem(parseInt(item.id))}>Delete</button>
                    </td>
                </tr>)}
            </tbody>
            </table>}
            {items.length === 0 && 
                <div>
                    No items to display! Get started by adding an item to your inventory.
            </div>}
        </div>
    )
};