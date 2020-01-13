package pg_astro

import "errors"

type SimpleProfile struct {
	D1 int `json:"d1"`
	P1 int `json:"p1"`
	D2 int `json:"d2"`
	P2 int `json:"p2"`
}

type ComplexProfile struct {
	Pwm0  int    `json:"pwm0"`
	Time1 string `json:"time1"`
	Pwm1  int    `json:"pwm1"`
	Time2 string `json:"time2"`
	Pwm2  int    `json:"pwm2"`
	Time3 string `json:"time3"`
	Pwm3  int    `json:"pwm3"`
	Time4 string `json:"time4"`
	Pwm4  int    `json:"pwm4"`
}

type Lamp struct {
	Dir             int                `json:"dir"`
	Level           int                `json:"level"`
	Nid             int                `json:"nid"`
	Group           int                `json:"group"`
	Mac             string             `json:"mac"`
	Smac            string             `json:"smac"`
	Rssi            int                `json:"rssi"`
	Devt            int                `json:"devt"`
	Devm            int                `json:"devm"`
	Eblk            *int               `json:"eblk,omitempty"`
	Cycles          *string            `json:"cycles,omitempty"`
	Runh            *int               `json:"runh,omitempty"`
	Nvsc            *int               `json:"nvsc,omitempty"`
	Lpwm            *int               `json:"lpwm,omitempty"`
	Cpwm            *int               `json:"cpwm,omitempty"`
	Mrssi           int                `json:"mrssi"`
	Rfch            *int               `json:"rfch,omitempty"`
	Rfpwr           *int               `json:"rfpwr,omitempty"`
	Pwm             *int               `json:"pwm,omitempty"`
	Pwmct           *int               `json:"pwmct,omitempty"`
	Pow             *int               `json:"pow,omitempty"`
	Lux             *int               `json:"lux,omitempty"`
	Temp            *int               `json:"temp,omitempty"`
	Energy          *int               `json:"energy,omitempty"`
	Rng             *int               `json:"rng,omitempty"`
	Tlevel          int                `json:"tlevel"`
	Date            *int64             `json:"date,omitempty"`
	Lat             *float64           `json:"lat,omitempty"`
	Lng             *float64           `json:"lon,omitempty"`
	Val             *int               `json:"val,omitempty"`
	Rise            *string            `json:"rise,omitempty"`
	Set             *string            `json:"set,omitempty"`
	Id              *int               `json:"id,omitempty"`
	SimpleProfiles  [3]*SimpleProfile  `json:"simpleProfiles"`
	ComplexProfiles [3]*ComplexProfile `json:"complexProfiles"`
	Scdtm           *int               `json:"scdtm,omitempty"`
	Rfps            *int               `json:"rfps,omitempty"`
	Twil            *int               `json:"twil,omitempty"`
}

func (l *Lamp) Apply(p *Package) {
	dir, ok := (*p)["dir"]
	if ok {
		l.Dir = int(dir.(float64))
	}

	level, ok := (*p)["level"]
	if ok {
		l.Level = int(level.(float64))
	}

	nid, ok := (*p)["nid"]
	if ok {
		l.Nid = int(nid.(float64))
	}

	group, ok := (*p)["group"]
	if ok {
		l.Group = int(group.(float64))
	}

	mac, ok := (*p)["mac"]
	if ok {
		l.Mac = mac.(string)
	}

	smac, ok := (*p)["smac"]
	if ok {
		l.Smac = smac.(string)
	}

	rssi, ok := (*p)["rssi"]
	if ok {
		l.Rssi = int(rssi.(float64))
	}

	devt, ok := (*p)["devt"]
	if ok {
		l.Devt = int(devt.(float64))
	}

	devm, ok := (*p)["devm"]
	if ok {
		l.Devm = int(devm.(float64))
	}

	eblk, ok := (*p)["eblk"]
	if ok {
		l.Eblk = new(int)
		*l.Eblk = int(eblk.(float64))
	}

	cycles, ok := (*p)["cycles"]
	if ok {
		l.Cycles = new(string)
		*l.Cycles = cycles.(string)
	}

	runh, ok := (*p)["runh"]
	if ok {
		l.Runh = new(int)
		*l.Runh = int(runh.(float64))
	}

	nsvc, ok := (*p)["nsvc"]
	if ok {
		l.Nvsc = new(int)
		*l.Nvsc = int(nsvc.(float64))
	}

	lpwm, ok := (*p)["lpwm"]
	if ok {
		l.Lpwm = new(int)
		*l.Lpwm = int(lpwm.(float64))
	}

	cpwm, ok := (*p)["cpwm"]
	if ok {
		l.Cpwm = new(int)
		*l.Cpwm = int(cpwm.(float64))
	}

	mrssi, ok := (*p)["mrssi"]
	if ok {
		l.Mrssi = int(mrssi.(float64))
	}

	rfch, ok := (*p)["rfch"]
	if ok {
		l.Rfch = new(int)
		*l.Rfch = int(rfch.(float64))
	}

	rfpwr, ok := (*p)["rfpwr"]
	if ok {
		l.Rfpwr = new(int)
		*l.Rfpwr = int(rfpwr.(float64))
	}

	pwm, ok := (*p)["pwm"]
	if ok {
		l.Pwm = new(int)
		*l.Pwm = int(pwm.(float64))
	}

	pwmct, ok := (*p)["pwmct"]
	if ok {
		l.Pwmct = new(int)
		*l.Pwmct = int(pwmct.(float64))
	}

	pow, ok := (*p)["pow"]
	if ok {
		l.Pow = new(int)
		*l.Pow = int(pow.(float64))
	}

	lux, ok := (*p)["lux"]
	if ok {
		l.Lux = new(int)
		*l.Lux = int(lux.(float64))
	}

	temp, ok := (*p)["temp"]
	if ok {
		l.Temp = new(int)
		*l.Temp = int(temp.(float64))
	}

	energy, ok := (*p)["energy"]
	if ok {
		l.Energy = new(int)
		*l.Energy = int(energy.(float64))
	}

	rng, ok := (*p)["rng"]
	if ok {
		l.Rng = new(int)
		*l.Rng = int(rng.(float64))
	}

	tlevel, ok := (*p)["tlevel"]
	if ok {
		l.Tlevel = int(tlevel.(float64))
	}

	date, ok := (*p)["date"]
	if ok {
		l.Date = new(int64)
		*l.Date = int64(date.(float64))
	}

	lat, ok := (*p)["lat"]
	if ok {
		l.Lat = new(float64)
		*l.Lat = lat.(float64)
	}

	lon, ok := (*p)["lon"]
	if ok {
		l.Lng = new(float64)
		*l.Lng = lon.(float64)
	}

	val, ok := (*p)["val"]
	if ok {
		l.Val = new(int)
		*l.Val = int(val.(float64))
	}

	rise, ok := (*p)["rise"]
	if ok {
		l.Rise = new(string)
		*l.Rise = rise.(string)
	}

	set, ok := (*p)["set"]
	if ok {
		l.Set = new(string)
		*l.Set = set.(string)
	}

	id, ok := (*p)["id"]
	if ok {
		l.Id = new(int)
		*l.Id = int(id.(float64))

		if *l.Id < 4 {
			l.updateSimpleProfile(*l.Id, p)
		} else {
			l.updateComplexProfile(*l.Id, p)
		}
	}

	scdtm, ok := (*p)["scdtm"]
	if ok {
		l.Scdtm = new(int)
		*l.Scdtm = int(scdtm.(float64))
	}

	rfps, ok := (*p)["rfps"]
	if ok {
		l.Rfps = new(int)
		*l.Rfps = int(rfps.(float64))
	}

	twil, ok := (*p)["twil"]
	if ok {
		l.Twil = new(int)
		*l.Twil = int(twil.(float64))
	}
}

func (l *Lamp) updateSimpleProfile(id int, p *Package) error {
	if l.SimpleProfiles[id-1] == nil {
		l.SimpleProfiles[id-1] = new(SimpleProfile)
	}

	d1, ok := (*p)["d1"]
	if !ok {
		return errors.New("d1 parameter not found")
	}
	l.SimpleProfiles[id-1].D1 = int(d1.(float64))

	p1, ok := (*p)["p1"]
	if !ok {
		return errors.New("p1 parameter not found")
	}
	l.SimpleProfiles[id-1].P1 = int(p1.(float64))

	d2, ok := (*p)["d2"]
	if !ok {
		return errors.New("d2 parameter not found")
	}
	l.SimpleProfiles[id-1].D2 = int(d2.(float64))

	p2, ok := (*p)["p2"]
	if !ok {
		return errors.New("p2 parameter not found")
	}
	l.SimpleProfiles[id-1].P2 = int(p2.(float64))

	return nil
}

func (l *Lamp) updateComplexProfile(id int, p *Package) error {
	if l.ComplexProfiles[id-4] == nil {
		l.ComplexProfiles[id-4] = new(ComplexProfile)
	}

	pwm0, ok := (*p)["pwm0"]
	if !ok {
		return errors.New("pwm0 parameter not found")
	}
	l.ComplexProfiles[id-4].Pwm0 = int(pwm0.(float64))

	time1, ok := (*p)["time1"]
	if !ok {
		return errors.New("time1 parameter not found")
	}
	l.ComplexProfiles[id-4].Time1 = time1.(string)

	pwm1, ok := (*p)["pwm1"]
	if !ok {
		return errors.New("pwm1 parameter not found")
	}
	l.ComplexProfiles[id-4].Pwm1 = int(pwm1.(float64))

	time2, ok := (*p)["time2"]
	if !ok {
		return errors.New("time2 parameter not found")
	}
	l.ComplexProfiles[id-4].Time2 = time2.(string)

	pwm2, ok := (*p)["pwm2"]
	if !ok {
		return errors.New("pwm2 parameter not found")
	}
	l.ComplexProfiles[id-4].Pwm2 = int(pwm2.(float64))

	time3, ok := (*p)["time3"]
	if !ok {
		return errors.New("time3 parameter not found")
	}
	l.ComplexProfiles[id-4].Time3 = time3.(string)

	pwm3, ok := (*p)["pwm3"]
	if !ok {
		return errors.New("pwm3 parameter not found")
	}
	l.ComplexProfiles[id-4].Pwm3 = int(pwm3.(float64))

	time4, _ := (*p)["time4"]
	if !ok {
		return errors.New("time4 parameter not found")
	}
	l.ComplexProfiles[id-4].Time4 = time4.(string)

	pwm4, _ := (*p)["pwm4"]
	if !ok {
		return errors.New("pwm4 parameter not found")
	}
	l.ComplexProfiles[id-4].Pwm4 = int(pwm4.(float64))

	return nil
}
