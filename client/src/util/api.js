import axios from 'axios';

axios.defaults.baseURL = 'http://localhost:8080';


// add the item to inventory
const AddItem = (item) => new Promise((resolve, reject) => {
    axios.post('/api/item', item)
        .then((response) => {
            resolve(response.data.item);
        })
        .catch((error) => {
            reject(error);
        });
});

// retrieves the list of items from the server
const GetItems = () => new Promise((resolve, reject) => {
    axios.get('/api/item')
        .then((response) => {
            resolve(response.data.items);
        })
        .catch((error) => {
            reject(error);
        });
});


// deletes the item from inventory
const DeleteItem = (itemId) => new Promise((resolve, reject) => {
    axios.delete(`/api/item/${itemId}`)
        .then( _ => {
            resolve();
        })
        .catch((error) => {
            reject(error);
        });
});

// updates the item in inventory
const UpdateItem = (item) => new Promise((resolve, reject) => {
    let itemId = item.id;
    delete item['id'];
    axios.patch(`/api/item/${itemId}`, item)
        .then((response) => {
            resolve(response.data.item);
        })
        .catch((error) => {
            reject(error);
        });
});


const DownloadCSV = () => new Promise((resolve, reject) => {
    axios.get('/api/item/csv')
        .then((response) => {
            resolve(response.data);
        })
        .catch((error) => {
            reject(error);
        });
});

export {
    AddItem,
    GetItems,
    DeleteItem,
    UpdateItem,
    DownloadCSV,
};
