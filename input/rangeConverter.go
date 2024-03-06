package input


type converter struct {
	minimumInput float64
	maximumImput float64

	minimumOutput float64
	maximumOutput float64
}


func (c *converter) convert(invalue float64) float64 {
	if invalue < c.minimumInput {
		invalue = c.minimumInput
	} else if invalue > c.maximumImput {
		invalue = c.maximumImput
	}

	interpolationFactor := ( invalue - c.minimumInput ) / ( c.maximumImput - c.minimumInput) 
	return (interpolationFactor * (c.maximumOutput - c.minimumOutput) + c.minimumOutput)
}

type RangeCoverter struct {
    
}

func (r *RangeCoverter) Convert(rangeid Range, invalue float64) float64 {
	return 0
}

// // Internal type shortcuts
// private:
// 	typedef std::map<Range, Converter> ConversionMapT;

// // Construction
// public:
// 	explicit RangeConverter(std::wifstream& infile);

// // Conversion interface
// public:
// 	template <typename RangeType>
// 	RangeType Convert(Range rangeid, RangeType invalue) const
// 	{
// 		ConversionMapT::const_iterator iter = ConversionMap.find(rangeid);
// 		if(iter == ConversionMap.end())
// 			return invalue;

// 		return iter->second.Convert<RangeType>(invalue);
// 	}

// // Internal tracking
// private:
// 	ConversionMapT ConversionMap;
// };