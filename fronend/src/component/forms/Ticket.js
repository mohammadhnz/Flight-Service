import {useForm, Controller} from "react-hook-form"
import DatePicker from "react-multi-date-picker";
import persian from "react-date-object/calendars/persian"
import persian_fa from "react-date-object/locales/persian_fa"
import Button from '@material-ui/core/Button';
import {Autocomplete, FormControlLabel, IconButton, Radio, RadioGroup, TextField} from "@mui/material";
// import countryNames from "../../static/countries.json";
import React, {useState, useEffect} from 'react';
// import flight_api from '../../api/flights'
import axios from "axios";


function Ticket() {
    const {register, handleSubmit, control, formState: {errors}} = useForm({mode: "onBlur",});
    const [submittedDep, setLeftDate] = useState();
    const [submittedRet, setReturnDate] = useState();
    const [ratioVal, setRatioVal] = useState('business');
    const [ratioValStatus, setRatioValStatus] = useState('around');
    const [passNum, setPassNum] = useState(0);
    const [options, setOptions] = useState([])

    // const options = countryNames.flights;

    const getSourceAndDestinationCountries = async () => {
        try {
            const response = await axios({
                method: 'get',
                url: 'http://localhost:8080/flights',
                headers: {
                    'Content-Type': 'application/json;charset=UTF-8',
                    'Access-Control-Allow-Origin': '*',
                },
            })
            setOptions(response.data)
            console.log(response.data)
        } catch
            (err) {
            console.log(err.response.data)
        }
    }
    useEffect(() => {
        getSourceAndDestinationCountries()
    }, [])
    const defaultProps = {
        options: options,
        getOptionLabel: (option) => option.city + "-" + option.name + "/" + option.airport + " " + option.iata,
    };


    const upPassNum = (event) => {
        if (!/[0-9]/.test(event.key)) {
            event.preventDefault();
            setPassNum(prevCount => parseInt(prevCount) + 1);
        }
    };

    const downPassNum = (event) => {
        if (!/[0-9]/.test(event.key)) {
            event.preventDefault();
            setPassNum(prevCount => parseInt(prevCount) - 1);
        }
    };

    const handlePasNumChange = (event) => {
        if (!/[0-9]/.test(event.key)) {
            event.preventDefault();
            setPassNum(event.target.value)
        }
    }
    const handleRatioChange = (event) => {
        setRatioVal(event.target.value);
    };

    const handleRaitoStatusChange = (event) => {
        setRatioValStatus(event.target.value)
    }

    function onSubmitButton(data, {retDate, leftDate}) {
        setLeftDate(leftDate);
        setReturnDate(retDate);
        data.passNum = passNum;
        let formData = new FormData();
        const retDateData = {
            "returnDay": data.retDate.day,
            "returnMonth": data.retDate.month.name,
            "returnYear": data.retDate.year

        }
        const leftDateData = {
            "returnDay": data.leftDate.day,
            "returnMonth": data.leftDate.month.name,
            "returnYear": data.leftDate.year

        }

        const fromId = data.from.split(" ").at(-1)
        const toId = data.to.split(" ").at(-1)
        formData.append("to", toId)
        formData.append("from", fromId)
        formData.append("class_name", data.class_name)
        formData.append("passNum", data.passNum)
        formData.append("return_status", ratioValStatus)
        formData.append("leftDate", JSON.stringify(leftDateData))
        formData.append("retDate", JSON.stringify(retDateData))
        for (const value of formData.values()) {
            console.log(value);
        }
        let object = {};
        formData.forEach(function (value, key) {
            object[key] = value;
        });
        let json = JSON.stringify(object);
        console.log(json)
        axios.post(`http://localhost:8888/tickets`, object)
            .then(res => {
                console.log(res);
                console.log(res.data);
            })
    }

    return (
        <>
            <form className="data-form classes.root" onSubmit={handleSubmit(onSubmitButton)}>
                <label htmlFor="from-field">
                    از
                    <Autocomplete
                        {...defaultProps}
                        autoHighlight
                        renderInput={(params) => (
                            <TextField {...params} {...register("from")} placeholder="from" id="from-field"
                                       variant="standard"/>
                        )}
                    />
                </label>
                <label htmlFor="to-field">
                    به
                    <Autocomplete
                        {...defaultProps}
                        autoHighlight
                        renderInput={(params) => (
                            <TextField {...params} {...register("to")} placeholder="to" id="to-field"
                                       variant="standard"/>
                        )}
                    />
                </label>

                <RadioGroup row aria-label="way" name="return_status" value={ratioValStatus}
                            onChange={handleRaitoStatusChange}>
                    <FormControlLabel value="around"
                                      control={<Radio id="field-around" {...register("return_status")}/>}
                                      label="رفت و برگشت"/>
                    <FormControlLabel value="oneway"
                                      control={<Radio id="field-oneway" {...register("return_status")}/>}
                                      label="یک طرفه"/>
                </RadioGroup>

                <Controller
                    control={control}
                    name="retDate"
                    rules={{required: true}} //optional
                    render={({
                                 field: {onChange, name, value},
                                 fieldState: {invalid, isDirty}, //optional
                                 formState: {errors}, //optional, but necessary if you want to show an error message
                             }) => (
                        <>
                            <DatePicker
                                value={value || ""}
                                onChange={(date) => {
                                    onChange(date?.isValid ? date : "");
                                }}
                                calendar={persian}
                                locale={persian_fa}
                                calendarPosition="bottom-right"
                                format={"YYYY/MM/DD"}

                                containerStyle={{
                                    marginLeft: '1.5rem'
                                }}
                            />
                            {errors && errors[name] && errors[name].type === "required" && (
                                //if you want to show an error message
                                <span>your error message !</span>
                            )}
                        </>
                    )}
                />
                <Controller
                    control={control}
                    name="leftDate"
                    rules={{required: true}} //optional
                    render={({
                                 field: {onChange, name, value},
                                 fieldState: {invalid, isDirty}, //optional
                                 formState: {errors}, //optional, but necessary if you want to show an error message
                             }) => (
                        <>
                            <DatePicker
                                value={value || ""}
                                onChange={(date) => {
                                    onChange(date?.isValid ? date : "");
                                }}
                                calendar={persian}
                                locale={persian_fa}
                                calendarPosition="bottom-right"
                                format={"YYYY/MM/DD"}
                            />
                            {errors && errors[name] && errors[name].type === "required" && (
                                //if you want to show an error message
                                <span>your error message !</span>
                            )}
                        </>
                    )}
                />
                <div>
                    <label htmlFor="passNum">تعداد مسافران</label>
                    <div style={{display: 'inline-block'}}>
                        <IconButton aria-label="plus" size="large" onClick={upPassNum}>+</IconButton>
                        <input style={{width: '25ch',}}
                               variant="standard" {...register("passNum", {required: true})}
                               value={passNum}
                               onChange={handlePasNumChange}/>
                        <IconButton aria-label="plus" size="large" onClick={downPassNum}>-</IconButton>
                    </div>
                </div>
                <br/>
                <RadioGroup row aria-label="classNames" name="class_name" value={ratioVal}
                            onChange={handleRatioChange}>
                    <FormControlLabel value="buisiness"
                                      control={<Radio id="field-around" {...register("class_name")}/>}
                                      label="بیزینس"/>
                    <FormControlLabel value="economy"
                                      control={<Radio id="field-oneway" {...register("class_name")}/>}
                                      label="اکونومی"/>
                    <FormControlLabel value="firstClass"
                                      control={<Radio id="field-oneway" {...register("class_name")}/>}
                                      label="فرست کلاس"/>
                </RadioGroup>
                <br/>
                <Button variant="outlined" type="submit" className="btn btn-primary">
                    Send
                </Button>
            </form>
        </>
    )
}


export default Ticket;