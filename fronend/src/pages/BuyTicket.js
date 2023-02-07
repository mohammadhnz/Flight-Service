import DrawerAppBar from "../component/pageElements/DrawerAppBar";
import {useLocation, useParams} from 'react-router-dom'
import Box from '@mui/material/Box';
import Card from '@mui/material/Card';
import CardActions from '@mui/material/CardActions';
import CardContent from '@mui/material/CardContent';
import Button from '@mui/material/Button';
import Typography from '@mui/material/Typography';
import React, {useEffect} from 'react'
import {yupResolver} from '@hookform/resolvers/yup'
import * as Yup from 'yup'
import {useForm, useFieldArray} from 'react-hook-form'
import Grid from "@mui/material/Grid";
import Ticket from "../component/forms/Ticket";

function BuyTicket() {
    const location = useLocation()
    const {from, to, flight_id, duration_min, duration_hour} = location.state

    const JsSchema = Yup.object().shape({
        FavFood: Yup.string().required('Value is mendatory!'),
        passenger: Yup.array().of(
            Yup.object().shape({
                name: Yup.string().required('Value is mendatory'),
            }),
        ),
    })
    const optionsDf = {resolver: yupResolver(JsSchema)}
    const {
        control,
        formState,
        handleSubmit,
        register,
        watch,
        reset,
    } = useForm(optionsDf)
    const {errors} = formState
    const {fields, append, remove} = useFieldArray({name: 'passenger', control})
    const FavFood = watch('FavFood')
    useEffect(() => {
        const currentProp = parseInt(FavFood || 0)
        const previousProp = fields.length
        if (currentProp > previousProp) {
            for (let i = previousProp; i < currentProp; i++) {
                append({name: ''})
            }
        } else {
            for (let i = previousProp; i > currentProp; i--) {
                remove(i - 1)
            }
        }
    }, [FavFood])

    function onSubmit(res) {
        const numberOfPassengers = parseInt(FavFood || 0)
        let formData = new FormData();
        let listOfPass = []
        for (let i = 0; i < numberOfPassengers; i++) {
            let passengerData = {
                "name": res.passenger[i].name,
                "family": res.lastname[i].name,
                "passport": res.pass[i].name
            }
            listOfPass.push(passengerData)
            console.log(passengerData)
        }
        listOfPass.forEach(item => {
            formData.append(`passenger`, JSON.stringify(item));
        });
        console.log(formData.getAll('passenger'));
    }

    const card = (
            <CardContent>
                <h2>اطلاعات مسافرین</h2>
                <Typography sx={{fontSize: 14}} color="text.secondary" gutterBottom>
                    {duration_hour} : {duration_min}
                </Typography>
                <Typography variant="h5" component="div">
                    از {from} به {to}
                </Typography>
            </CardContent>
    );
    return (
        <div className=" rmdp-rtl">
            <Grid justifyContent="center">
                <Card>{card}</Card>
            </Grid>
            <Grid container justifyContent="center">
                <div className="container mt-5">
                    <form onSubmit={handleSubmit(onSubmit)}>
                        <div className="form-group">
                            <label className="mb-2">تعداد مسافرین:</label>
                            <select
                                name="FavFood"
                                {...register('FavFood')}
                                className={`form-control ${errors.FavFood ? 'is-invalid' : ''}`}
                            >
                                {['Select Options', 1, 2, 3, 4, 5, 6].map((i) => (
                                    <option key={i} value={i}>
                                        {i}
                                    </option>
                                ))}
                            </select>
                            <div className="invalid-feedback">{errors.FavFood?.message}</div>
                        </div>
                        {fields.map((item, i) => (
                            <div key={i} className="mt-3 mb-2">
                                <div>
                                    <strong className="text-primary">passenger {i + 1}</strong>
                                    <div className="form-group">
                                        <input
                                            name={`passenger[${i}]name`}
                                            {...register(`passenger.${i}.name`)}
                                            className={`form-control ${
                                                errors.passenger?.[i]?.name ? 'is-invalid' : ''
                                            }`}
                                            type="text"
                                        />
                                        <input
                                            name={`lastname[${i}]name`}
                                            {...register(`lastname.${i}.name`)}
                                            className={`form-control ${
                                                errors.lastname?.[i]?.name ? 'is-invalid' : ''
                                            }`}
                                            type="text"
                                        />
                                        <input
                                            name={`pass[${i}]name`}
                                            {...register(`pass.${i}.name`)}
                                            className={`form-control ${
                                                errors.pass?.[i]?.name ? 'is-invalid' : ''
                                            }`}
                                            type="text"
                                        />
                                        <div className="invalid-feedback">
                                            {errors.passenger?.[i]?.name?.message}
                                        </div>
                                    </div>
                                </div>
                            </div>
                        ))}
                        <button type="submit" className="btn btn-success mt-3 mb-2">
                            ثبت
                        </button>
                        <button
                            onClick={() => reset()}
                            type="button"
                            className="btn btn-info">
                            ریست کردن انتخاب ها
                        </button>
                    </form>
                </div>
            </Grid>
        </div>

    )
}

export default BuyTicket