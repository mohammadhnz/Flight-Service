import {Dashboard} from '../component/pageElements/Dashboard';
import {dataProvider} from '../api/DataProvider';
import {Admin, Resource} from "react-admin";
import jsonServerProvider from "ra-data-json-server";

export default function AdminPanel() {
    // const dataProvider = jsonServerProvider("https://jsonplaceholder.typicode.com");
    return (
        <>
            <Admin dataProvider={dataProvider} dashboard={Dashboard}>
                <Resource name="flights"/>
            </Admin>
        </>
    )
}