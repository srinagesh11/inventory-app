import { useState, useEffect } from 'react';
import { AddItem, GetItems, DeleteItem, UpdateItem, DownloadCSV } from '../util/api';
import ItemsToolbar from './ItemsToolbar';
import ItemsList from './itemsList';
import Alert from './Alert';

export default function ItemsDashboard(props) {

    const [items, setItems] = useState([]);
    const [alert, setAlert] = useState({ message: "", type: "" });
    const [selectedItem, setSelectedItem] = useState(null);

    const setErr = (message) => setAlert({ message: message, type: "danger" });

 
    useEffect(() => {
        // call the API to get the list of items
        GetItems()
            .then((items) => {
                setItems(items);
            })
            .catch((error) => {
                console.log(error);
                setErr("There was an error retrieving the list of items");
            });
    }, []);

    // add the item to inventory
    const onAddItem = (item, cb) => {
        setSelectedItem(null);
        AddItem(item).then((item) => {
            setItems(prev => [...prev, item]);
            setAlert({ message: "Item added successfully", type: "success" });
            cb(true);
        }).catch(error => {
            console.log(error);
            if (error.response) {
                setErr(error.response.data.error);
            } else {
                setErr("An error occurred while adding the item to inventory");
            }
        });
    };

    // handler for selecting an item from the list
    const onSelectItemToEdit = (item) => {
        setSelectedItem(item);
    }

    // handler for editing an item
    const onEditItem = (item, cb) => { 
        setSelectedItem(null);
        UpdateItem(item).then((item) => {
            setItems(prev => prev.map(i => i.id === item.id ? item : i));
            setAlert({ message: "Item updated successfully", type: "success" });
            cb(true);
        }).catch(error => {
            console.log(error);
            if (error.response) {
                setErr(error.response.data.error);
            } else {
                setErr("An error occurred while updating the item");
            }
        });
    };

    // handler for the delete button
    const onDeleteItem = (itemId) => {
        let confirm = window.confirm("Are you sure you want to delete this item?");
        if (!confirm) return;
        DeleteItem(itemId).then(_ => {
            setItems(prev => prev.filter(i => i.id !== itemId));
            setAlert({ message: "Item deleted successfully", type: "success" });
        }).catch(error => {
            console.log(error);
            if (error.response) {
                setErr(error.response.data.error);
            } else {
                setErr("An error occurred while deleting the item from inventory");
            }
        })
    };

    // download the CSV file
    const downloadCSV = () => {
        DownloadCSV().then(csv => {
            let content = "data:text/csv;charset=utf-8," + csv;
            let encodedUri = encodeURI(content);
            let link = document.createElement("a");
            link.setAttribute("href", encodedUri);
            link.setAttribute("download", "inventory.csv");
            document.body.appendChild(link); 
            link.click();
            document.body.removeChild(link);
        }).catch(err => {
            console.log(err);
            if (err.response) {
                setErr(err.response.data.error);
            } else {
                setErr("An error occurred while downloading the CSV file");
            }
        });
    };
    

    return (
        <div>
            <Alert message={alert.message} type={alert.type} onClose={event => setAlert({message: "", type: ""})}  />
            <ItemsToolbar
                onAddItem={onAddItem}
                setErr={setErr}
                selectedItem={selectedItem}
                onEditItem={onEditItem}
                downloadCSV={downloadCSV}
            />
            <ItemsList 
                items={items}
                onSelectItemToEdit={onSelectItemToEdit}
                onDeleteItem={onDeleteItem}
            />
        </div>
    )
};
