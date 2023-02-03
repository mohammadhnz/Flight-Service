import {useForm, Controller} from "react-hook-form"
import {useState} from "react";
import DatePicker from "react-multi-date-picker";
import persian from "react-date-object/calendars/persian"
import persian_fa from "react-date-object/locales/persian_fa"
import Button from '@material-ui/core/Button';
import {Autocomplete, FormControlLabel, Radio, RadioGroup, TextField} from "@mui/material";
import {makeStyles} from "@material-ui/core/styles";
import countryNames from "../../static/countries.json";
import InputNumber from 'rmc-input-number';


const useStyles = makeStyles((theme) => ({
    root: {
        display: 'flex',
        flexWrap: 'wrap',
    },
    input: {
        marginLeft: theme.spacing(1),
        marginRight: theme.spacing(1),
        width: '25ch',
    },
}));

function Ticket() {
    const {register, handleSubmit, control, formState: {errors}} = useForm({mode: "onBlur",});
    const [submittedDep, setLeftDate] = useState();
    const [submittedRet, setReturnDate] = useState();
    const [ratioVal, setRatioVal] = useState('around');
    const classes = useStyles();
    const options = countryNames;
    const defaultProps = {
        options: options,
        getOptionLabel: (option) => option.name,
    };


    const handleRatioChange = (event) => {
        setRatioVal(event.target.value);
    };

    function onSubmitButton(data, {retDate, leftDate}) {
        setLeftDate(leftDate);
        setReturnDate(retDate);
        console.log(data)
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

                <RadioGroup row aria-label="way" name="way" value={ratioVal} onChange={handleRatioChange}>
                    <FormControlLabel value="around" control={<Radio id="field-around" {...register("return-status")}/>}
                                      label="رفت و برگشت"/>
                    <FormControlLabel value="oneway" control={<Radio id="field-oneway" {...register("return-status")}/>}
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

                <br/>
                <InputNumber {...register("passengers")} />
                <br/>
                <Button variant="outlined" type="submit" className="btn btn-primary">
                    Send
                </Button>
            </form>
        </>
    )
}


export default Ticket;