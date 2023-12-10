<?php

namespace App\Http\Requests;

class ClassSubjectMaterialAttachmentRequest extends FormRequest
{
    /**
     * Determine if the user is authorized to make this request.
     */
    public function authorize(): bool
    {
        return true;
    }

    /**
     * Get the validation rules that apply to the request.
     *
     * @return array<string, \Illuminate\Contracts\Validation\ValidationRule|array<mixed>|string>
     */
    public function rules(): array
    {
        return [
            'class_subject_material_id' => 'required|numeric|gt:0',
            'content_type' => 'required',
            'content' => 'required|string',
        ];
    }

    /**
     * @return array
     * Custom validation message
     */
    public function messages(): array
    {
        return [
            'class_subject_material_id.required' => 'The class subject material ID field is required.',
            'class_subject_material_id.numeric' => 'The class subject material ID must be a number.',
            'class_subject_material_id.gt' => 'The class subject material ID must be a positive number.',
            'content_type.required' => 'The content type field is required.',
            'content.required' => 'The content field is required.',
        ];
    }
}
